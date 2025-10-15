package service 

import  (
	"context"
	"github.com/nullsec45/golang-anime-restapi/domain"
	"github.com/nullsec45/golang-anime-restapi/internal/config"
	"github.com/nullsec45/golang-anime-restapi/dto"
	"github.com/nullsec45/golang-anime-restapi/internal/utility"
	// "github.com/nullsec45/golang-anime-restapi/internal/cache"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"database/sql"
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
		return dto.AuthResponse{}, domain.AuthFail
	}
	
	err = utility.VerifyPassword(user.Password, req.Password)
	if err != nil {
		return dto.AuthResponse{}, domain.AuthFail
	}

	claim := jwt.MapClaims {
		"id" : user.Id,
		"exp" : time.Now().Add(time.Duration(auth.config.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(auth.config.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, domain.AuthFail
	}

	return dto.AuthResponse {
		Token:tokenStr,
		UserID:user.Id,
		Email:user.Email,
	}, nil
}

func (auth AuthService) Register (ctx context.Context, req dto.RegisterRequest) error {
	userData, err := auth.userRepository.FindByEmail(ctx, req.Email)

	if err != nil {
		return err
	}
	
	if userData.Email != "" {
		return domain.EmailRegister
	}
	
	password := req.Password
	confirmPassword := req.ConfirmPassword

	if password != confirmPassword {
		return domain.PasswordNotMatch
	}
	
	hashPassword, err := utility.HashPassword(req.ConfirmPassword)

	if err != nil {
		return err
	}


	user := domain.User {
		Id:uuid.New().String(),
		Email:req.Email,
		Password:hashPassword,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return auth.userRepository.Save(ctx, &user)
}