package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"github.com/nullsec45/golang-anime-restapi/internal/session"
	"github.com/sirupsen/logrus"
	"errors"
)

type AuthApi struct {
	authService domain.AuthService
	session *session.Manager
}

func NewAuth(app *fiber.App, authService domain.AuthService, session *session.Manager) {
	api := AuthApi{
		authService: authService,
		session:session,
	}

	auth := app.Group("/auth")
	auth.Post("/login", api.Login)
	auth.Post("/register", api.Register)
	auth.Put("/update-password", api.UpdatePassword)
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
			"fails": fails,
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

		if errors.Is(err, domain.ErrAuthFailed) {
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
			"ua": ctx.Get("User-Agent"),
		})

		return ctx.Status(statusCode).JSON(
			dto.CreateResponseError(
				statusCode,
				err.Error(),
			),
		)
	}

	if err := api.session.Renew(ctx); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, "Session renew failed"))
	}

	if err := api.session.SetUser(ctx, res.UserID, res.Email); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, "Session set failed"))
	}


	utility.CreateLog("info", "login success", "activity", logrus.Fields{
		"route": "/v1/auth/login",
		"email": utility.MaskEmail(req.Email),
		"ip":     ctx.IP(),
		"ua": ctx.Get("User-Agent"),
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
		utility.CreateLog("warn", "register failed: invalid request body", "application", logrus.Fields{
			"route": "/v1/auth/register",
			"reason": "invalid_body",
			"error": err.Error(),
			"ip": ctx.IP(),
			"ua": ctx.Get("User-Agent"),
		})


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
		utility.CreateLog("warn", "register failed: validation error", "application", logrus.Fields{
				"route": "/v1/auth/register",
				"reason": "validation_failed",
				"fails": fails,
				"ip": ctx.IP(),
				"ua": ctx.Get("User-Agent"),
		})
		
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed register account",
				fails,
			),
		)
	}

	err := api.authService.Register(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, domain.EmailRegister) {
			statusCode = 409

			utility.CreateLog("info", "register failed: invalid register", "activity", logrus.Fields{
				"route": "/v1/auth/register",
				"email": utility.MaskEmail(req.Email),
				"reason":"email_already_register",
				"ip":    ctx.IP(),
				"ua":    ctx.Get("User-Agent"),
			})
    	}

		utility.CreateLog("error", "register failed: internal error", "application", logrus.Fields{
			"route": "/v1/auth/register",
			"error": err.Error(),
			"ip":    ctx.Context().RemoteAddr().String(),
			"ua":    ctx.Get("User-Agent"),
		})

		return ctx.Status(statusCode).JSON(
			dto.CreateResponseError(
				statusCode,
				err.Error(),
			),
		)
	}

	utility.CreateLog("info", "register success", "activity", logrus.Fields{
		"route": "/v1/auth/register",
		"email": utility.MaskEmail(req.Email),
		"ip":     ctx.IP(),
		"ua": ctx.Get("User-Agent"),
	})

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Register."),
	)
}

func (api AuthApi) UpdatePassword(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	_, email, err := api.session.GetUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.CreateResponseError(fiber.StatusUnauthorized, "Unauthorized"))
	}

	var req dto.UpdatePasswordRequest
	if err := ctx.BodyParser(&req); err != nil {
		utility.CreateLog("warn", "update password failed: invalid request body", "application", logrus.Fields{
			"route": "/v1/auth/update-password",
			"reason": "invalid_body",
			"error": err.Error(),
			"ip": ctx.IP(),
			"ua": ctx.Get("User-Agent"),
		})


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
		utility.CreateLog("warn", "update password failed: validation error", "application", logrus.Fields{
			"route": "/v1/auth/update-password",
			"reason": "validation_failed",
			"fails": fails,
			"ip": ctx.IP(),
			"ua": ctx.Get("User-Agent"),
		})
		
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(
				http.StatusBadRequest,
				"Failed update password",
				fails,
			),
		)
	}

	err = api.authService.UpdatePassword(c, req, email)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, domain.CurrentPasswordWrong) {
			statusCode = 409

			utility.CreateLog("info", "update password failed: invalid password", "activity", logrus.Fields{
				"route": "/v1/auth/update-password",
				"email": utility.MaskEmail(email),
				"reason":"current_password_wrong",
				"ip":    ctx.IP(),
				"ua":    ctx.Get("User-Agent"),
			})
    	}

		utility.CreateLog("error", "update password failed: internal error", "application", logrus.Fields{
			"route": "/v1/auth/update-password",
			"error": err.Error(),
			"ip":    ctx.Context().RemoteAddr().String(),
			"ua":    ctx.Get("User-Agent"),
		})

		return ctx.Status(statusCode).JSON(
			dto.CreateResponseError(
				statusCode,
				err.Error(),
			),
		)
	}

	utility.CreateLog("info", "update password success", "activity", logrus.Fields{
		"route": "/v1/auth/update-password",
		"email": utility.MaskEmail(email),
		"ip":     ctx.IP(),
		"ua": ctx.Get("User-Agent"),
	})

	return ctx.Status(http.StatusOK).JSON(
		dto.CreateResponseSuccess("Successfully Update Password."),
	)
}