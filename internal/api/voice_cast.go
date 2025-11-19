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

type VoiceCastAPI struct {
	voiceCastService domain.VoiceCastService
}

func NewVoiceCast(
	app *fiber.App, 
	voiceCastService domain.VoiceCastService,
	authMiddleware fiber.Handler,
) {
	vcAPI := VoiceCastAPI{
		voiceCastService: voiceCastService,
	}

	voiceCast := app.Group("/voice-cast", authMiddleware)

	voiceCast.Post("/", vcAPI.Create)
	voiceCast.Put(":id", vcAPI.Update)
	// voiceCast.Delete("anime/:animeId", vcAPI.DeleteByAnimeId)
	// voiceCast.Delete("tag/:tagId", vcAPI.DeleteByTagId)
	voiceCast.Delete(":id", vcAPI.DeleteById)
}

func (vca VoiceCastAPI) Create (ctx *fiber.Ctx) error {
	vct, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateVoiceCastRequest

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

	err := vca.voiceCastService.Create(vct, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (vca VoiceCastAPI) Update (ctx *fiber.Ctx) error {
	vct, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateVoiceCastRequest

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
	err := vca.voiceCastService.Update(vct, req)
	
	statusCode := http.StatusInternalServerError
	
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}


// func (vca VoiceCastAPI) DeleteByAnimeId (ctx *fiber.Ctx) error {
// 	vct, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
// 	defer cancel()

// 	id := ctx.Params("animeId")
// 	err := vca.voiceCastService.DeleteByAnimeId(vct, id)

// 	statusCode := http.StatusInternalServerError

// 	if err != nil {
// 		if errors.Is(err, domain.ErrNotFound) {
// 			statusCode = http.StatusNotFound
// 		}
// 		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
// }

// func (vca VoiceCastAPI) DeleteByTagId (ctx *fiber.Ctx) error {
// 	vct, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
// 	defer cancel()

// 	id := ctx.Params("tagId")
// 	err := vca.voiceCastService.DeleteByTagId(vct, id)

// 	statusCode := http.StatusInternalServerError

// 	if err != nil {
// 		if errors.Is(err, domain.ErrNotFound) {
// 			statusCode = http.StatusNotFound
// 		}
// 		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
// 	}

	
// 	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
// }

func (vca VoiceCastAPI) DeleteById (ctx *fiber.Ctx) error {
	vct, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := vca.voiceCastService.DeleteById(vct, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}