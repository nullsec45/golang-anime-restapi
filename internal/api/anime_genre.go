package api

import (
	"net/http"
	"time"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"errors"
)

type AnimeGenresAPI struct {
	animeGenresService domain.AnimeGenresService
}

func NewAnimeGenres(
	app *fiber.App, 
	animeGenresService domain.AnimeGenresService,
	authMiddleware fiber.Handler,
) {
	agAPI := AnimeGenresAPI{
		animeGenresService: animeGenresService,
	}

	animeGenres := app.Group("/anime-genres", authMiddleware)

	animeGenres.Post("/", agAPI.Create)
	animeGenres.Put(":id", agAPI.Update)
	animeGenres.Delete("anime/:animeId", agAPI.DeleteByAnimeId)
	animeGenres.Delete("genre/:genreId", agAPI.DeleteByGenreId)
	animeGenres.Delete(":id", agAPI.DeleteById)
}

func (aga AnimeGenresAPI) Create (ctx *fiber.Ctx) error {
	ang, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeGenresRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest, 
				"Failed created data", map[string]string{
				"body": err.Error(),
			}),
		)
	}
	
	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			http.StatusBadRequest,
			"Failed created data",
			fails,
		))
	}

	err := aga.animeGenresService.Create(ang, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(
			http.StatusInternalServerError,
			err.Error(),
		))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (aga AnimeGenresAPI) Update (ctx *fiber.Ctx) error {
	ang, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAnimeGenresRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed updated data",
				map[string]string{
					"body": err.Error(),
				},
			),
		)
	}

	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			http.StatusBadRequest,
			"Failed updated data",
			fails,
		))
	}

	req.Id=ctx.Params("id")
	err := aga.animeGenresService.Update(ang, req)

	if errors.Is(err, domain.AnimeGenresNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}


func (aga AnimeGenresAPI) DeleteByAnimeId (ctx *fiber.Ctx) error {
	ang, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("animeId")
	err := aga.animeGenresService.DeleteByAnimeId(ang, id)

	if errors.Is(err, domain.AnimeGenresNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (aga AnimeGenresAPI) DeleteByGenreId (ctx *fiber.Ctx) error {
	ang, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("genreId")
	err := aga.animeGenresService.DeleteByGenreId(ang, id)

	if errors.Is(err, domain.AnimeGenresNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(
			http.StatusInternalServerError,
			err.Error(),
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (aga AnimeGenresAPI) DeleteById (ctx *fiber.Ctx) error {
	ang, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := aga.animeGenresService.DeleteById(ang, id)

	if errors.Is(err, domain.AnimeGenresNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(
			http.StatusInternalServerError,
			err.Error(),
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}