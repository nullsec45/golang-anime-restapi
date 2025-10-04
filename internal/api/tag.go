package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"time"
	"net/http"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"errors"
)

type AnimeTagAPI struct {
	animeTagService domain.AnimeTagService
}

func NewAnimeTag(
	app * fiber.App, 
	animeTagService domain.AnimeTagService,
	authMiddleware fiber.Handler,
) {
	anmTA := AnimeTagAPI{
		animeTagService: animeTagService,
	}

	customer := app.Group("/tags", authMiddleware)

	customer.Get("", anmTA.Index)
	customer.Post("", anmTA.Create)
	customer.Put(":id", anmTA.Update)
	customer.Delete(":id", anmTA.Delete)
	customer.Get(":id", anmTA.Show)
}

func (anmTA AnimeTagAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := anmTA.animeTagService.Index(c)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data",res))
}

func (anmTA AnimeTagAPI) Show (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := anmTA.animeTagService.Show(c, id)
	
	if errors.Is(err, domain.AnimeTagNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError,err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
}

func (anmTA AnimeTagAPI) Create (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeTagRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(http.StatusBadRequest,"Failed created data", map[string]string{
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

	err := anmTA.animeTagService.Create(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError,err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (anmTA AnimeTagAPI) Update (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAnimeTagRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(http.StatusBadRequest,"Failed updated data", map[string]string{
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
	err := anmTA.animeTagService.Update(c,req)

	if errors.Is(err, domain.AnimeTagNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError,err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}

func (anmTA AnimeTagAPI) Delete (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := anmTA.animeTagService.Delete(c, id)

	if errors.Is(err, domain.AnimeTagNotFound) {
        return ctx.Status(http.StatusNotFound).JSON(dto.CreateResponseError(http.StatusNotFound, err.Error()))
    }

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError,err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}