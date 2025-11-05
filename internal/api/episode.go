package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"errors"
)

type EpisodeAPI struct {
	animeEpisodeService domain.AnimeEpisodeService
}

func NewAnimeEpisode(
	app *fiber.App,
	animeEpisodeService domain.AnimeEpisodeService,
	authMiddleware fiber.Handler,
) {
	epAPI := EpisodeAPI{
		animeEpisodeService: animeEpisodeService,
	}

	episode := app.Group("/episodes", authMiddleware)
	episode.Post("/", epAPI.Create)
	episode.Patch(":id", epAPI.Update)
	episode.Delete("anime/:animeId", epAPI.DeleteByAnimeId)
	episode.Delete(":id", epAPI.DeleteById)
	episode.Get("anime/:animeId", epAPI.Index)
	episode.Get(":id", epAPI.Show)
}

func (epa EpisodeAPI) Index(ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var q dto.PaginationQuery
	if err := ctx.QueryParser(&q); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(http.StatusBadRequest, "invalid query params: " + err.Error()))
	} 

	q.Normalize(1, 10, 100)

	epts := domain.EpisodeListOptions{
		Pagination: q,
		Filter: domain.EpisodeFilter{
			Search: q.Search,
		},
	}

	animeId := ctx.Params("animeId")

	res, err := epa.animeEpisodeService.Index(an, animeId, epts)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(200).JSON(dto.CreateResponseSuccessWithDataPagination("Successfully Get Data", res))
}

func (epa  EpisodeAPI) Show (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := epa.animeEpisodeService.Show(an, id)
  	
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
}

func (epa EpisodeAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeEpisodeRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed created data",
				map[string]string{"body": err.Error()},
			),
		)
	}

	if err := req.Validate(); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(
			dto.CreateResponseErrorData(
				http.StatusUnprocessableEntity,
				"Validation failed",
				map[string]string{"body": err.Error()},
			),
		)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(
			dto.CreateResponseErrorData(
				http.StatusUnprocessableEntity,
				"Validation failed",
				fails,
			),
		)
	}

	if err := epa.animeEpisodeService.Create(c, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			dto.CreateResponseError(
				http.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(http.StatusCreated).JSON(
		dto.CreateResponseSuccess("Successfully created data."),
	)
}

func (epa EpisodeAPI) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseError(http.StatusBadRequest, "id is required"),
		)
	}

	var req dto.UpdateAnimeEpisodeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed updated data",
				map[string]string{"body": err.Error()},
			),
		)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(
			dto.CreateResponseErrorData(
				http.StatusUnprocessableEntity,
				"Validation failed",
				fails,
			),
		)
	}

	req.Id = id
	err := epa.animeEpisodeService.Update(c, req)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Updated Data"),
	)
}

func (epa EpisodeAPI) DeleteByAnimeId(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("animeId")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseError(
				http.StatusBadRequest,
				"animeId is required",
			),
		)
	}

	err := epa.animeEpisodeService.DeleteByAnimeId(c, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Deleted Data"),
	)
}

func (epa EpisodeAPI) DeleteById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseError(
				http.StatusBadRequest,
				"id is required",
			),
		)
	}

	err := epa.animeEpisodeService.DeleteById(c, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Deleted Data"),
	)
}
