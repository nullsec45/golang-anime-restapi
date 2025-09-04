package domain

import (
	"context"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}