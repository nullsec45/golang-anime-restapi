package domain

import "errors"

var (
	ErrNotFound          = errors.New("Not Found!")
	ErrAlreadyExist      = errors.New("Already Exist!")
	ErrAuthFailed        = errors.New("Authentication Failed")
	ErrForbidden         = errors.New("forbidden")
	ErrValidation        = errors.New("validation failed")
 	EmailRegister        = errors.New("Email already register!.")
	CurrentPasswordWrong = errors.New("Current Password is Wrong!.")
	AnimeMediaOutsideDir = errors.New("Invalid media path: outside of base dir")
)



