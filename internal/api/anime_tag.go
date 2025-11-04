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

type AnimeTagsAPI struct {
	animeTagsService domain.AnimeTagsService
}

func NewAnimeTags(
	app *fiber.App, 
	animeTagsService domain.AnimeTagsService,
	authMiddleware fiber.Handler,
) {
	atsAPI := AnimeTagsAPI{
		animeTagsService: animeTagsService,
	}

	animeTags := app.Group("/anime-tags", authMiddleware)

	animeTags.Post("/", atsAPI.Create)
	animeTags.Put(":id", atsAPI.Update)
	animeTags.Delete("anime/:animeId", atsAPI.DeleteByAnimeId)
	animeTags.Delete("tag/:tagId", atsAPI.DeleteByTagId)
	animeTags.Delete(":id", atsAPI.DeleteById)
}

func (ata AnimeTagsAPI) Create (ctx *fiber.Ctx) error {
	ant, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeTagsRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest, 
				"Failed created data", 
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
			"Failed created data",
			fails,
		))
	}

	err := ata.animeTagsService.Create(ant, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (ata AnimeTagsAPI) Update (ctx *fiber.Ctx) error {
	ant, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAnimeTagsRequest

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
	err := ata.animeTagsService.Update(ant, req)
	
	statusCode := http.StatusInternalServerError
	
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}


func (ata AnimeTagsAPI) DeleteByAnimeId (ctx *fiber.Ctx) error {
	ant, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("animeId")
	err := ata.animeTagsService.DeleteByAnimeId(ant, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (ata AnimeTagsAPI) DeleteByTagId (ctx *fiber.Ctx) error {
	ant, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("tagId")
	err := ata.animeTagsService.DeleteByTagId(ant, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (ata AnimeTagsAPI) DeleteById (ctx *fiber.Ctx) error {
	ant, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ata.animeTagsService.DeleteById(ant, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}