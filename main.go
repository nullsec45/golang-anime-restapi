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
	"log"
	// "github.com/nullsec45/golang-anime-restapi/internal/cache"
	"github.com/nullsec45/golang-anime-restapi/internal/session"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"fmt"
)

func main(){
	conf := config.Get()
	dbConnection := connection.GetDatabase(conf.Database)
	rdb, err := connection.GetRedisClient(conf.Redis)

	if err != nil { 
		utility.CreateLog("warn", fmt.Sprintf("failed connect redis: %v", err), "application")
		log.Fatalf("failed connect redis: %v", err)
	}

	defer rdb.Close()

	
	sessionStorage, err := connection.GetRedisStorage(conf.Redis)

	if err != nil {
		utility.CreateLog("warn", fmt.Sprintf("Failed connect to redis: %v", err), "application")
		log.Fatalf("failed init redis storage: %v", err)
	}

	defer connection.CloseRedisStorage(sessionStorage)

	secureCookie := conf.App.AppEnv == "production"
	sessions := session.NewWithRedisStorage(sessionStorage, secureCookie)

	app := fiber.New(
		fiber.Config{
			ProxyHeader: fiber.HeaderXForwardedFor,
			EnableTrustedProxyCheck: true,
			TrustedProxies: []string{
				"10.0.0.0/8",	
				"172.16.0.0/12",
				"192.168.0.0/16",
			},
		},
	)


	authMiddleware := jwtMid.New(
		jwtMid.Config{
			SigningKey:jwtMid.SigningKey{Key:[]byte(conf.Jwt.Key)},
			ErrorHandler:func (ctx *fiber.Ctx, err error) error {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError(http.StatusUnauthorized, "Unauthenticated, please login!."))
			},
		},
	)

	userRepository := repository.NewUser(dbConnection)
	animeRepository := repository.NewAnime(dbConnection)
	animeEpisodeRepository := repository.NewAnimeEpisode(dbConnection)
	animeGenreRepository := repository.NewAnimeGenre(dbConnection)
	animeGenresRepository := repository.NewAnimeGenres(dbConnection)
	animeTagRepository := repository.NewAnimeTag(dbConnection)
	animeTagsRepository := repository.NewAnimeTags(dbConnection)
	mediaRepository := repository.NewMedia(dbConnection)
	animeStudioRepository := repository.NewAnimeStudio(dbConnection)
	animeStudiosRepository := repository.NewAnimeStudios(dbConnection)
	peopleRepository := repository.NewPeople(dbConnection)
	characterRepository := repository.NewCharacter(dbConnection)

	authService := service.NewAuth(conf, userRepository)
	animeService := service.NewAnime(conf, animeRepository, animeEpisodeRepository, animeGenreRepository, animeTagRepository, mediaRepository, animeStudioRepository)
	animeEpisodeService := service.NewAnimeEpisode(animeRepository, animeEpisodeRepository, mediaRepository, conf)
	animeGenreService := service.NewAnimeGenre(animeGenreRepository)
	animeGenresService := service.NewAnimeGenres(animeRepository, animeGenreRepository, animeGenresRepository)
	animeTagService :=  service.NewAnimeTag(animeTagRepository)
	animeTagsService := service.NewAnimeTags(animeRepository, animeTagRepository, animeTagsRepository)
	mediaService := service.NewMedia(conf, mediaRepository)
	animeStudioService := service.NewAnimeStudio(animeStudioRepository)
	animeStudiosService := service.NewAnimeStudios(animeRepository, animeStudioRepository, animeStudiosRepository)
	peopleService := service.NewPeople(peopleRepository)
	characterService := service.NewCharacter(characterRepository)

	v1 := fiber.New()
	api.NewAuth(v1, authService, sessions)
	api.NewAnime(v1, animeService, authMiddleware)
	api.NewAnimeEpisode(v1, animeEpisodeService, authMiddleware)
	api.NewAnimeGenre(v1, animeGenreService, authMiddleware)
	api.NewAnimeGenres(v1, animeGenresService, authMiddleware)
	api.NewAnimeTag(v1, animeTagService, authMiddleware)
	api.NewAnimeTags(v1, animeTagsService, authMiddleware)
	api.NewMedia(app, conf, mediaService, authMiddleware)
	api.NewAnimeStudio(v1, animeStudioService, authMiddleware)
	api.NewAnimeStudios(v1, animeStudiosService, authMiddleware)
	api.NewPeople(v1, peopleService, authMiddleware)
	api.NewCharacter(v1, characterService, authMiddleware)
	
	app.Mount("/v1", v1)

	_ = app.Listen(conf.Server.Host +":"+ conf.Server.Port)
}