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
	"errors"
	// "encoding/json"
)

type CharacterAPI struct {
	characterService domain.CharacterService
}

func NewCharacter(
	app *fiber.App, 
	characterService domain.CharacterService,
	authMiddleware fiber.Handler,
) {
	ca := CharacterAPI{
		characterService: characterService,
	}

	character := app.Group("/characters", authMiddleware)

	character.Get("", ca.Index)
	character.Post("", ca.Create)
	character.Put(":id", ca.Update)
	character.Delete(":id", ca.Delete)
	character.Get(":id", ca.Show)
}

func (ca CharacterAPI) Index(ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var q dto.PaginationQuery
	if err := ctx.QueryParser(&q); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(http.StatusBadRequest, "invalid query params: " + err.Error()))
	} 

	q.Normalize(1, 10, 100)

	opts := domain.CharacterListOptions{
		Pagination: q,
		Filter: domain.CharacterFilter{
			Search: q.Search,
		},
	}

	res, err := ca.characterService.Index(an, opts)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(200).JSON(dto.CreateResponseSuccessWithDataPagination("Successfully Get Data", res))
}

func (ca CharacterAPI) Show (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := ca.characterService.Show(an, id)
  	
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
}

func (ca CharacterAPI) Create (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCharacterRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(http.StatusBadRequest, "Failed created data", map[string]string{
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

	err := ca.characterService.Create(an, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (ca CharacterAPI) Update (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCharacterRequest

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
	err := ca.characterService.Update(an,req)
		
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}

func (ca CharacterAPI) Delete (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ca.characterService.Delete(an, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}
