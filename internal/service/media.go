package service

import (
	"context"
	"time"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"github.com/google/uuid"
	"database/sql"
	"path"
	"path/filepath"
	"errors"
	"os"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	"fmt"
)

type MediaService struct {
	config *config.Config
	mediaRepository domain.MediaRepository
}

func NewMedia(config *config.Config, mediaRepository domain.MediaRepository) domain.MediaService {
	return &MediaService{
		config:config,
		mediaRepository:mediaRepository,
	}
}

func (m MediaService) Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
	media := domain.Media{
		Id:uuid.NewString(),
		Path:req.Path,
		CreatedAt:sql.NullTime{Time:time.Now(), Valid:true},
	}

	err := m.mediaRepository.Save(ctx, &media)
	if err != nil {
		return dto.MediaData{}, err
	}

	url := path.Join(m.config.Server.Asset, media.Path)
	return dto.MediaData{
		Id:media.Id,
		Path:media.Path,
		Url:url,
	}, nil
}

func (m MediaService) Show (ctx context.Context, id string) (dto.MediaData, error) {
	exist, err :=  m.mediaRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return dto.MediaData{}, domain.AnimeMediaNotFound
    }
    
    if err != nil {
        return dto.MediaData{}, err
    }

    return dto.MediaData{
		Id:exist.Id,
		Path:exist.Path,
	}, nil
}

func (m MediaService) View(ctx context.Context, id string) (absPath, filename string, modTime time.Time, err error) {
	fmt.Println(id)
	media, err := m.mediaRepository.FindById(ctx, id)
	fmt.Println(media)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || media.Id == "" {
			return "", "", time.Time{}, domain.AnimeMediaNotFound
		}
		return "", "", time.Time{}, err
	}
	if media.Id == "" {
		return "", "", time.Time{}, domain.AnimeMediaNotFound
	}

	absFile, err := utility.SafeJoin(m.config.Storage.BasePath, media.Path)
	if err != nil {
		return "", "", time.Time{}, err
	}

	st, statErr := os.Stat(absFile)
	if statErr != nil {
		if errors.Is(statErr, os.ErrNotExist) {
			return "", "", time.Time{}, domain.AnimeMediaNotFound
		}
		return "", "", time.Time{}, statErr
	}

	return absFile, filepath.Base(media.Path), st.ModTime(), nil
}

func (m MediaService) Update(ctx context.Context, req dto.UpdateMediaRequest) (dto.MediaData, error) {
	exist, err := m.mediaRepository.FindById(ctx, req.Id)

	oldPath := exist.Path

    if err != nil && exist.Id == "" {
        return dto.MediaData{}, domain.AnimeMediaNotFound
    }

	fmt.Println(req.Path)

	exist.Path = req.Path	
	exist.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	media, err := m.mediaRepository.Update(ctx, &exist)

	if err != nil {
		return dto.MediaData{}, err
	}

	url := path.Join(m.config.Server.Asset, media.Path)
	return dto.MediaData{
		Id:exist.Id,
		Path:media.Path,
		OldPath:oldPath,
		Url:url,
	}, nil
}

func (m MediaService) Delete (ctx context.Context, id string) (path string, err error) {
    exist, err := m.mediaRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
       return "", domain.AnimeMediaNotFound
    }
    
    if err != nil {
        return "", err
    }

	if err := m.mediaRepository.Delete(ctx, id); err != nil {
		return "", err
	}

    return exist.Path, nil
}