package main

import (	
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"github.com/nullsec45/golang-anime-restapi/internal/connection"
	"github.com/nullsec45/golang-anime-restapi/internal/repository"
	"github.com/nullsec45/golang-anime-restapi/internal/service"
	"github.com/nullsec45/golang-anime-restapi/internal/api"
	"github.com/nullsec45/golang-anime-restapi/dto"
	jwtMid "github.com/gofiber/contrib/jwt"
	"net/http"
)

func main(){
	conf := config.Get()
	dbConnection := connection.GetDatabase(conf.Database)

	app := fiber.New()

	authMiddleware := jwtMid.New(
		jwtMid.Config{
			SigningKey:jwtMid.SigningKey{Key:[]byte(conf.Jwt.Key)},
			ErrorHandler:func (ctx *fiber.Ctx, err error) error {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("Endpoint perlu token, silahkan login terlebih dahulu."))
			},
		},
	)

	userRepository := repository.NewUser(dbConnection)
	animeRepository := repository.NewAnime(dbConnection)


	authService := service.NewAuth(conf, userRepository)
	animeService := service.NewAnime(animeRepository)

	api.NewAuth(app, authService)
	api.NewAnime(app, animeService, authMiddleware)

	_ = app.Listen(conf.Server.Host +":"+ conf.Server.Port)
}