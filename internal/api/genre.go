package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"time"
	"net/http"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
)

type AnimeGenreAPI struct {
	animeGenreService domain.AnimeGenreService
}

func NewAnimeGenre(
	app * fiber.App, 
	animeGenreService domain.AnimeGenreService,
	authMiddleware fiber.Handler,
) {
	anmGA := AnimeGenreAPI{
		animeGenreService: animeGenreService,
	}

	customer := app.Group("/genres", authMiddleware)

	customer.Get("", anmGA.Index)
	customer.Post("", anmGA.Create)
	customer.Put(":id", anmGA.Update)
	customer.Delete(":id", anmGA.Delete)
	customer.Get(":id", anmGA.Show)
}

func (anmGA AnimeGenreAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := anmGA.animeGenreService.Index(c)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data",res))
}

func (anmGA AnimeGenreAPI) Create (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeGenreRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData("Failed created data", map[string]string{
				"body": err.Error(),
			}),
		)
	}

	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			"Failed created data",
			fails,
		))
	}

	err := anmGA.animeGenreService.Create(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (anmGA AnimeGenreAPI) Update (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAnimeGenreRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData("Failed updated data", map[string]string{
				"body": err.Error(),
			}),
		)
	}

	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			"Failed updated data",
			fails,
		))
	}

	req.Id=ctx.Params("id")
	err := anmGA.animeGenreService.Update(c,req)
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}

func (anmGA AnimeGenreAPI) Delete (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := anmGA.animeGenreService.Delete(c, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (anmGA AnimeGenreAPI) Show (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := anmGA.animeGenreService.Show(c, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
}