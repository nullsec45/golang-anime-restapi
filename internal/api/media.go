package api 

import (
	"net/http"
	"context"
	"time"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/google/uuid"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
)

type MediaAPI struct {
	config *config.Config
	mediaService domain.MediaService
}

func NewMedia(
	app *fiber.App,
	config *config.Config,
	mediaService domain.MediaService,
	authMiddleware fiber.Handler,
) {
	ma := MediaAPI {
		config:config,
		mediaService:mediaService,
	}

	app.Static("/assets", config.Storage.BasePath)

	app.Post("/media", authMiddleware, ma.Create)
	app.Delete("/media/:id", authMiddleware, ma.Delete)

}

func (ma MediaAPI) Create (ctx *fiber.Ctx) error {
	if ma.mediaService == nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError("mediaService is not initialized"))
	}

	if ma.config == nil || ma.config.Storage.BasePath == "" {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError("storage config is not initialized"))
	}

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	filename := uuid.NewString() + file.Filename
	path := filepath.Join(ma.config.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	res, err :=	ma.mediaService.Create(c, dto.CreateMediaRequest{
		Path:filename,
	}) 

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccessWithData("Successfullly Create Media",res))
}

func (ma MediaAPI) Delete (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ma.mediaService.Delete(an, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Media"))
}
