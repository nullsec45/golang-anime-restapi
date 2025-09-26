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
				return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("Unauthenticated, please login!."))
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

	authService := service.NewAuth(conf, userRepository)
	animeService := service.NewAnime(animeRepository, animeEpisodeRepository, animeGenreRepository, animeTagRepository)
	animeEpisodeService := service.NewAnimeEpisode(animeRepository, animeEpisodeRepository)
	animeGenreService := service.NewAnimeGenre(animeGenreRepository)
	animeGenresService := service.NewAnimeGenres(animeRepository, animeGenreRepository, animeGenresRepository)
	animeTagService :=  service.NewAnimeTag(animeTagRepository)
	animeTagsService := service.NewAnimeTags(animeRepository, animeTagRepository, animeTagsRepository)
	mediaService := service.NewMedia(conf, mediaRepository)

	v1 := fiber.New()
	api.NewAuth(v1, authService)
	api.NewAnime(v1, animeService, authMiddleware)
	api.NewAnimeEpisode(v1, animeEpisodeService, authMiddleware)
	api.NewAnimeGenre(v1, animeGenreService, authMiddleware)
	api.NewAnimeGenres(v1, animeGenresService, authMiddleware)
	api.NewAnimeTag(v1, animeTagService, authMiddleware)
	api.NewAnimeTags(v1, animeTagsService, authMiddleware)
	api.NewMedia(app, conf, mediaService, authMiddleware)
	
	app.Mount("/v1", v1)

	_ = app.Listen(conf.Server.Host +":"+ conf.Server.Port)
}