package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
)

type AuthApi struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	api := AuthApi{
		authService: authService,
	}

	auth := app.Group("/auth")
	auth.Post("/login", api.Login)
	auth.Post("/register", api.Register)
}

func (api AuthApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed created data",
				map[string]string{"body": err.Error()},
			),
		)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed created data",
				fails,
			),
		)
	}

	res, err := api.authService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			dto.CreateResponseError(
				http.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccessWithData("Successfully Login.", res),
	)
}

func (api AuthApi) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed Register",
				map[string]string{"body": err.Error()},
			),
		)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed created data",
				fails,
			),
		)
	}

	// NOTE: jangan shadowing err (pakai =, bukan :=)
	err := api.authService.Register(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			dto.CreateResponseError(
				http.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Register."),
	)
}
