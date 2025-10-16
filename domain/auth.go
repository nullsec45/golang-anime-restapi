package domain

import (
	"context"
	"github.com/nullsec45/golang-anime-restapi/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest)  error
	UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest, email string) error
}