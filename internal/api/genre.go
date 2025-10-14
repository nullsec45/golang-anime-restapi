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

type AnimeGenreAPI struct {
	animeGenreService domain.AnimeGenreService
}

func NewAnimeGenre(
	app *fiber.App,
	animeGenreService domain.AnimeGenreService,
	authMiddleware fiber.Handler,
) {
	anmGA := AnimeGenreAPI{
		animeGenreService: animeGenreService,
	}

	genre := app.Group("/genres", authMiddleware)

	genre.Get("", anmGA.Index)
	genre.Post("", anmGA.Create)
	genre.Put(":id", anmGA.Update)
	genre.Delete(":id", anmGA.Delete)
	genre.Get(":id", anmGA.Show)
}

func (anmGA AnimeGenreAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := anmGA.animeGenreService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			dto.CreateResponseError(http.StatusInternalServerError, err.Error()),
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccessWithData("Successfully Get Data", res),
	)
}

func (anmGA AnimeGenreAPI) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := anmGA.animeGenreService.Show(c, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.AnimeGenreNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccessWithData("Successfully Get Data", res),
	)
}

func (anmGA AnimeGenreAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeGenreRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed created data",
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

	if err := anmGA.animeGenreService.Create(c, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			dto.CreateResponseError(http.StatusInternalServerError, err.Error()),
		)
	}

	return ctx.Status(http.StatusCreated).JSON(
		dto.CreateResponseSuccess("Successfully created data."),
	)
}

func (anmGA AnimeGenreAPI) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseError(http.StatusBadRequest, "id is required"),
		)
	}

	var req dto.UpdateAnimeGenreRequest
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
	err := anmGA.animeGenreService.Update(c, req)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.AnimeGenreNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Updated Data"),
	)
}

func (anmGA AnimeGenreAPI) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseError(http.StatusBadRequest, "id is required"),
		)
	}

	err := anmGA.animeGenreService.Delete(c, id)
	
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.AnimeGenreNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Deleted Data"),
	)
}
