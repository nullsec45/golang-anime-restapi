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
)

type AuthApi struct {
	authService domain.AuthService
}

func NewAuth(
	app *fiber.App, 
	authService domain.AuthService,
) {
	api := AuthApi {
		authService : authService,
	}

	app.Post("/auth/login", api.Login)
	app.Post("/auth/register", api.Register)
}

func (api AuthApi) Login (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.AuthRequest
	
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

	res, err := api.authService.Login(c, req)
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully Login.", res))
}

func (api AuthApi) Register (ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.RegisterRequest 
	
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData("Failed Register", map[string]string{
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

	err := api.authService.Register(c, req)
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Register."))
}