package api

import (
	"net/http"
	"time"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"fmt"
	// "encoding/json"
)

type AnimeAPI struct {
	animeService domain.AnimeService
}

func NewAnime(
	app *fiber.App, 
	animeService domain.AnimeService,
	authMiddleware fiber.Handler,
) {
	anmApi := AnimeAPI{
		animeService: animeService,
	}

	anime := app.Group("/animes", authMiddleware)

	anime.Get("", anmApi.Index)
	anime.Post("", anmApi.Create)
	anime.Put(":id", anmApi.Update)
	anime.Delete(":id", anmApi.Delete)
	anime.Get(":id", anmApi.Show)
}

func (ana AnimeAPI) Index(ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var q dto.PaginationQuery
	if err := ctx.QueryParser(&q); err != nil {
	fmt.Println(q)
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("invalid query params: " + err.Error()))
	} 

	q.Normalize(1, 10, 100)

	opts := domain.AnimeListOptions{
		Pagination: q,
		Filter: domain.AnimeFilter{
			Search: q.Search,
		},
	}

	res, err := ana.animeService.Index(an, opts)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(200).JSON(dto.CreateResponseSuccessWithDataPagination("Successfully Get Data", res))
}

func (ana AnimeAPI) Show (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := ana.animeService.Show(an, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
}

func (ana AnimeAPI) Create (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeRequest

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

	err := ana.animeService.Create(an, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (ana AnimeAPI) Update (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateAnimeRequest

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
	err := ana.animeService.Update(an,req)
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}

func (ana AnimeAPI) Delete (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ana.animeService.Delete(an, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}
