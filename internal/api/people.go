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

type PeopleAPI struct {
	peopleService domain.PeopleService
}

func NewPeople(
	app *fiber.App, 
	peopleService domain.PeopleService,
	authMiddleware fiber.Handler,
) {
	peApi := PeopleAPI{
		peopleService: peopleService,
	}

	people := app.Group("/peoples", authMiddleware)

	people.Get("", peApi.Index)
	people.Post("", peApi.Create)
	people.Put(":id", peApi.Update)
	people.Delete(":id", peApi.Delete)
	people.Get(":id", peApi.Show)
}

func (pa PeopleAPI) Index(ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var q dto.PaginationQuery
	if err := ctx.QueryParser(&q); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(http.StatusBadRequest, "invalid query params: " + err.Error()))
	} 

	q.Normalize(1, 10, 100)

	opts := domain.PeopleListOptions{
		Pagination: q,
		Filter: domain.PeopleFilter{
			Search: q.Search,
		},
	}

	res, err := pa.peopleService.Index(an, opts)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(200).JSON(dto.CreateResponseSuccessWithDataPagination("Successfully Get Data", res))
}

func (pa PeopleAPI) Show (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	res, err := pa.peopleService.Show(an, id)
  	
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Get Data", res))
}

func (pa PeopleAPI) Create (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreatePeopleRequest

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

	err := pa.peopleService.Create(an, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("Successfully created data."))
}

func (pa PeopleAPI) Update (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdatePeopleRequest

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
	err := pa.peopleService.Update(an,req)
		
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Updated Data"))
}

func (pa PeopleAPI) Delete (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := pa.peopleService.Delete(an, id)

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Data"))
}
