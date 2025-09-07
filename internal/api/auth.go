package api 

import (
	"net/http"
	"time"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
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
}

func (api AuthApi) Login (ctx *fiber.Ctx) error {
	// fmt.Println("Auth Login Function Called")
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.AuthRequest 
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	res, err := api.authService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccessWithData("Successfully login.", res))
}