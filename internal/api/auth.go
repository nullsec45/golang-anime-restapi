package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"github.com/sirupsen/logrus"
	"errors"
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
		utility.CreateLog("warn", "login failed: invalid request body", "application", logrus.Fields{
			"route": "/v1/auth/login",
			"reason": "invalid_body",
			"error": err.Error(),
			"ip": ctx.IP(),
			"ua": ctx.Get("User-Agent"),
		})

		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed Login.",
				map[string]string{"body": err.Error()},
			),
		)
	}

		fails := utility.Validate(req)
		if len(fails) > 0 {
			utility.CreateLog("warn", "login failed: validation error", "application", logrus.Fields{
				"route": "/v1/auth/login",
				"reason": "validation_failed",
				"ip": ctx.IP(),
				"fails": fails, // Go1.21: atau kumpulkan manual
			})
			
			return ctx.Status(http.StatusBadRequest).JSON(
				dto.CreateResponseErrorData(
					http.StatusBadRequest,
					"Failed Login.",
					fails,
				),
			)
		}

	res, err := api.authService.Login(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, domain.AuthFail) {
			statusCode = 401

			utility.CreateLog("info", "login failed: invalid credentials", "activity", logrus.Fields{
				"route": "/v1/auth/login",
				"email": utility.MaskEmail(req.Email),
				"ip":    ctx.IP(),
				"ua":    ctx.Get("User-Agent"),
				"reason":"invalid_credentials",
			})
    	}

		utility.CreateLog("error", "login failed: internal error", "application", logrus.Fields{
			"route": "/v1/auth/login",
			"error": err.Error(),
			"ip":    ctx.Context().RemoteAddr().String(),
		})

		return ctx.Status(statusCode).JSON(
			dto.CreateResponseError(
				statusCode,
				err.Error(),
			),
		)
	}

	utility.CreateLog("info", "login success", "activity", logrus.Fields{
		"route": "/v1/auth/login",
		"email": utility.MaskEmail(req.Email),
		"ip":     ctx.IP(),
	})

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
