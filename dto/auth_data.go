package dto

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	UserID string `json:"-"`
	Email  string `json:"-"`
}

type RegisterRequest struct {
	Id                 string  `json:"id"`
	Email              string  `json:"email" validate:"required,email"`
	Password           string  `json:"password" validate:"required"`
	ConfirmPassword    string  `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type UpdatePasswordRequest struct {
	CurrentPassword    string  `json:"current_password" validate:"required"`
	NewPassword        string  `json:"new_password" validate:"required"`
	ConfirmPassword    string  `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}