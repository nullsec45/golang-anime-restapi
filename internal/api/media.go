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
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"errors"
	// "fmt"
	"os"
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

	media := app.Group("media", authMiddleware)

	media.Post("/", authMiddleware, ma.Create)
	media.Delete(":id", authMiddleware, ma.Delete)
	media.Get(":id",authMiddleware, ma.Get)
	media.Patch(":id", authMiddleware, ma.Update)

}

func (ma MediaAPI) Create (ctx *fiber.Ctx) error {

	if ma.mediaService == nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, "mediaService is not initialized"))
	}

	if ma.config == nil || ma.config.Storage.BasePath == "" {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(http.StatusInternalServerError, "storage config is not initialized"))
	}

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	allowed := []string{".jpg",".jpeg",".png",".mp4",".mkv",".avi",".flv"}
	const maxMB=20
	const maxBytes=maxMB * 1024 * 1024

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(http.StatusBadRequest,err.Error()))
		
	}

	if file == nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(http.StatusBadRequest, "Validation failed", map[string]string{
				"media": "File 'Media' files are required to be uploaded.",
			}),
		)
	}

	if vErr := utility.ValidateMediaFile(file, allowed, maxBytes); vErr != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(
			dto.CreateResponseErrorData(http.StatusUnprocessableEntity, "Validation failed", map[string]string{
				"media": vErr.Error(),
			}),
		)
	}

	ext := filepath.Ext(file.Filename)
	filename := uuid.NewString() + ext
	path := filepath.Join(ma.config.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	req := dto.CreateMediaRequest{	
		Path: filename,
	}

	res, err :=	ma.mediaService.Create(c, req) 

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			http.StatusBadRequest,
			"Failed uploaded media",
			fails,
		))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccessWithData("Successfullly Create Media",res))
}

func (ma MediaAPI) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	abs, filename, modTime, err := ma.mediaService.View(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return ctx.Status(fiber.StatusNotFound).
				JSON(dto.CreateResponseError(http.StatusNotFound, "Media not found"))
		}
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	ctx.Set(fiber.HeaderAcceptRanges, "bytes")
	ctx.Set(fiber.HeaderCacheControl, "public, max-age=31536000, immutable")
	ctx.Set(fiber.HeaderLastModified, modTime.UTC().Format(http.TimeFormat))
	ctx.Response().Header.SetCanonical([]byte("Content-Disposition"), []byte(`inline; filename="`+filename+`"`))

	ctx.Type(filepath.Ext(filename))

	return ctx.SendFile(abs, true) 
}

func (ma MediaAPI) Update (ctx *fiber.Ctx) error {

	if ma.mediaService == nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, "mediaService is not initialized"))
	}

	if ma.config == nil || ma.config.Storage.BasePath == "" {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(http.StatusInternalServerError, "storage config is not initialized"))
	}

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	allowed := []string{".jpg",".jpeg",".png",".mp4",".mkv",".avi",".flv"}
	const maxMB=20
	const maxBytes=maxMB * 1024 * 1024

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(http.StatusBadRequest,err.Error()))
		
	}

	if file == nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			dto.CreateResponseErrorData(http.StatusBadRequest, "Validation failed", map[string]string{
				"media": "File 'Media' files are required to be uploaded.",
			}),
		)
	}

	if vErr := utility.ValidateMediaFile(file, allowed, maxBytes); vErr != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(
			dto.CreateResponseErrorData(http.StatusUnprocessableEntity, "Validation failed", map[string]string{
				"media": vErr.Error(),
			}),
		)
	}

	ext := filepath.Ext(file.Filename)
	filename := uuid.NewString() + ext
	path := filepath.Join(ma.config.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(http.StatusInternalServerError, err.Error()))
	}

	req := dto.UpdateMediaRequest{	
		Path: filename,
	}

	id := ctx.Params("id")
	req.Id=id
	res, err :=	ma.mediaService.Update(c, req) 

	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	fails := utility.Validate(req)
	
	if len(fails) > 0{
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData(
			http.StatusBadRequest,
			"Failed uploaded media",
			fails,
		))
	}

	absFile, err := utility.SafeJoin(ma.config.Storage.BasePath, res.OldPath)
	
	if err != nil {
		return err
	}
	
	if rmErr := os.Remove(absFile); rmErr != nil && !os.IsNotExist(rmErr) {
		return rmErr
	}
	

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccessWithData("Successfullly Update Media",res))
}

func (ma MediaAPI) Delete (ctx *fiber.Ctx) error {
	an, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	id := ctx.Params("id")

	path, err := ma.mediaService.Delete(an, id)

	absFile, err := utility.SafeJoin(ma.config.Storage.BasePath, path)
	
	if err != nil {
		return err
	}
	
	if rmErr := os.Remove(absFile); rmErr != nil && !os.IsNotExist(rmErr) {
		return rmErr
	}
	
	statusCode := http.StatusInternalServerError

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}
		return ctx.Status(statusCode).JSON(dto.CreateResponseError(statusCode, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateResponseSuccess("Successfully Deleted Media"))
}
