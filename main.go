package main

import (	
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"github.com/nullsec45/golang-anime-restapi/internal/connection"
	"github.com/nullsec45/golang-anime-restapi/internal/repository"
	"github.com/nullsec45/golang-anime-restapi/internal/service"
	"github.com/nullsec45/golang-anime-restapi/internal/api"
	// jwtMid "github.com/gofiber/contrib/jwt"
	// "net/http"
	// "github.com/nullsec45/golang-anime-restapi/dto"
)

func main(){
	conf := config.Get()
	dbConnection := connection.GetDatabase(conf.Database)

	app := fiber.New()
	userRepository := repository.NewUser(dbConnection)


	authService := service.NewAuth(conf, userRepository)

	api.NewAuth(app, authService)

	_ = app.Listen(conf.Server.Host +":"+ conf.Server.Port)
}