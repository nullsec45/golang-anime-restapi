package main

import (	
	"github.com/gofiber/fiber/v2"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	// "github.com/nullsec45/golang-anime-restapi/internal/connection"
	// "github.com/nullsec45/golang-anime-restapi/internal/repository"
	// "github.com/nullsec45/golang-anime-restapi/internal/service"
	// "github.com/nullsec45/golang-anime-restapi/internal/api"
	// jwtMid "github.com/gofiber/contrib/jwt"
	// "net/http"
	// "github.com/nullsec45/golang-anime-restapi/dto"
	"fmt"
)

func main(){
	conf := config.Get()
	// dbConnection := connection.GetDatabase(conf.Database)
	fmt.Println("TEST")
	app := fiber.New()
	fmt.Println(conf.Server.Host)

	_ = app.Listen(conf.Server.Host +":"+ conf.Server.Port)
}