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

type AnimeStudiosAPI struct {
	animeStudiosService domain.AnimeStudiosService
}

func NewAnimeStudios(
	app *fiber.App, 
	animeStudiosService domain.AnimeStudiosService,
	authMiddleware fiber.Handler,
) {
	astAPI := AnimeStudiosAPI{
		animeStudiosService: animeStudiosService,
	}

	animeStudios := app.Group("/anime-studios", authMiddleware)

	animeStudios.Post("/", astAPI.Create)
	animeStudios.Put(":id", astAPI.Update)
	animeStudios.Delete("anime/:animeId", astAPI.DeleteByAnimeId)
	animeStudios.Delete("studio/:studioId", astAPI.DeleteByStudioId)
	animeStudios.Delete(":id", astAPI.DeleteById)
}

func (asa AnimeStudiosAPI) Create (ctx *fiber.Ctx) error {
	ans, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeStudiosRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed created data", 
				map[string]string{
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

	err := asa.animeStudiosService.Create(ans, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (asa AnimeStudiosAPI) Update (ctx *fiber.Ctx) error {
	ans, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAnimeStudiosRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(http.StatusBadRequest, "Failed updated data", map[string]string{
				"body": err.Error(),
			}),
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
	err := asa.animeStudiosService.Update(ans, req)

	statusCode := http.StatusInternalServerError
	
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}


	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}


func (asa AnimeStudiosAPI) DeleteByAnimeId (ctx *fiber.Ctx) error {
	ans, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("animeId")
	err := asa.animeStudiosService.DeleteByAnimeId(ans, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (asa AnimeStudiosAPI) DeleteByStudioId (ctx *fiber.Ctx) error {
	ans, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("studioId")
	err := asa.animeStudiosService.DeleteByStudioId(ans, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (asa AnimeStudiosAPI) DeleteById (ctx *fiber.Ctx) error {
	ans, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := asa.animeStudiosService.DeleteById(ans, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}