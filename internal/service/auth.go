package service 

import  (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	config *config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService{
	return AuthService {
		config:cnf,
		userRepository:userRepository,
	}
}

func (auth AuthService) Login (ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error){
	user, err := auth.userRepository.FindByEmail(ctx, req.Email)
	
	if err != nil {
		return dto.AuthResponse{}, err
	}
	
	if user.Id == "" {
		return dto.AuthResponse{}, errors.New("Autentikasi gagal")
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return dto.AuthResponse{},errors.New("Autentikasi gagal")
	}

	claim := jwt.MapClaims {
		"id" : user.Id,
		"exp" : time.Now().Add(time.Duration(auth.config.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(auth.config.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, errors.New("Autentikasi gagal")
	}

	return dto.AuthResponse {
		Token:tokenStr,
	}, nil
}