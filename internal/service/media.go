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
	"strings"
	"os"
	"errors"
	
)

type MediaService struct {
	config *config.Config
	mediaRepository domain.MediaRepository
}

func safeJoin(baseDir,rel string) (string, error) {
	cleanRel := filepath.Clean(rel)
	abs := filepath.Join(baseDir, cleanRel)

	baseAbs, err := filepath.Abs(baseDir)

	if err != nil {
		return "", err
	}
	absTarget, err := filepath.Abs(abs)
	if err != nil {
		return "", err
	}

	sep := string(filepath.Separator)
	if !strings.HasPrefix(absTarget+sep, baseAbs+sep) && absTarget != baseAbs {
		return "", domain.AnimeMediaOutsideDir
	}

	return absTarget, nil
}

func publicURL(baseURL, rel string) string {
	trimmed := strings.TrimRight(baseURL, "/")
	return path.Join(trimmed+"/", rel)
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


func (m MediaService) GetAbsPath(ctx context.Context, id string) (absPath, filename string, modTime time.Time, err error) {
	media, err := m.mediaRepository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || media.Id == "" {
			return "", "", time.Time{}, domain.AnimeMediaNotFound
		}
		return "", "", time.Time{}, err
	}
	if media.Id == "" {
		return "", "", time.Time{}, domain.AnimeMediaNotFound
	}

	absFile, err := safeJoin(m.config.Storage.BasePath, media.Path)
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

// func (m MediaService) Update(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
// 	media := domain.Media{
// 		Id:uuid.NewString(),
// 		Path:req.Path,
// 		CreatedAt:sql.NullTime{Time:time.Now(), Valid:true},
// 	}

// 	err := m.mediaRepository.Update(ctx, &media)
// 	if err != nil {
// 		return dto.MediaData{}, err
// 	}

// 	url := path.Join(m.config.Server.Asset, media.Path)
// 	return dto.MediaData{
// 		Id:media.Id,
// 		Path:media.Path,
// 		Url:url,
// 	}, nil
// }

func (m MediaService) Delete (ctx context.Context, id string) error {
    exist, err := m.mediaRepository.FindById(ctx, id)

    if err != nil && exist.Id == "" {
        return  domain.AnimeMediaNotFound
    }
    
    if err != nil {
        return err
    }

	absFile, err := safeJoin(m.config.Storage.BasePath, exist.Path)
	
	if err != nil {
		return err
	}
	
	if rmErr := os.Remove(absFile); rmErr != nil && !os.IsNotExist(rmErr) {
		return rmErr
	}

    return m.mediaRepository.Delete(ctx, id)
}