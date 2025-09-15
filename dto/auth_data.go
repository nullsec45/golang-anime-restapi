package dto

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Id                 string  `json:"id"`
	Email              string  `json:"email" validate:"required,email"`
	Password           string  `json:"password" validate:"required"`
	ConfirmPassword    string  `json:"confirm_password" validate:"required"`
}