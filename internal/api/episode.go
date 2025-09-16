package api

import (
	"net/http"
	"time"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	// "fmt"
	// "encoding/json"
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

	// episode.Get("", epAPI.Index)
	episode.Post("/", epAPI.Create)
	// episode.Put(":id", epAPI.Update)
	episode.Delete("anime/:animeId", epAPI.DeleteByAnimeId)
	episode.Delete(":id", epAPI.DeleteById)
	// episode.Get(":id", epAPI.Show)
}

// func (epa EpisodeAPI) Index(ctx *fiber.Ctx) error {
// 	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
// 	defer cancel()

// 	res, err := epa.animeEpisodeService.Index(an)

// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data",res))
// }

// func (epa EpisodeAPI) Show (ctx *fiber.Ctx) error {
// 	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
// 	defer cancel()

// 	id := ctx.Params("id")
// 	res, err := epa.animeEpisodeService.Show(an, id)

// 	if err != nil {
// 		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
// }

func (epa EpisodeAPI) Create (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateAnimeEpisodeRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData("Failed created data", map[string]string{
				"body": err.Error(),
			}),
		)
	}

	if err := req.Validate(); err != nil {
		// mapping validator errors -> 422 response, dll
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

	err := epa.animeEpisodeService.Create(an, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

// func (epa EpisodeAPI) Update (ctx *fiber.Ctx) error {
// 	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
// 	defer cancel()

// 	var req dto.UpdateEpisodeRequest

// 	if err := ctx.BodyParser(&req); err != nil {
// 		return ctx.SendStatus(http.StatusUnprocessableEntity)
// 	}
// 	fails := utility.Validate(req)
	
// 	if len(fails) > 0{
// 		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
// 			"Failed updated data",
// 			fails,
// 		))
// 	}

// 	req.Id=ctx.Params("id")
// 	err := epa.animeEpisodeService.Update(an,req)
	
// 	if err != nil {
// 		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
// }

func (epa EpisodeAPI) DeleteByAnimeId (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("animeId")
	err := epa.animeEpisodeService.DeleteByAnimeId(an, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}

func (epa EpisodeAPI) DeleteById (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := epa.animeEpisodeService.DeleteById(an, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}