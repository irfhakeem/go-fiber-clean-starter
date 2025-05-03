package dto

import "errors"

const (
	FAILED_UNAUTHORIZED    = "unauthorized"
	FAILED_FORBIDDEN       = "forbidden"
	FAILED_TOKEN_INVALID   = "token invalid"
	FAILED_TOKEN_NOT_FOUND = "token not found"
	FAILED_TOKEN_EXPIRED   = "token expired"
	FAILED_REGISTER_USER   = "failed to register user"
	FAILED_LOGIN_USER      = "failed to login user"
	FAILED_VERIFY_USER     = "failed to verify user"

	SUCCESS_REGISTER_USER = "success to register user"
	SUCCESS_LOGIN_USER    = "success to login user"
	SUCCESS_VERIFY_USER   = "success to verify user"
)

var (
	ErrGenerateToken      = errors.New("failed to create token")
	ErrCreateRefreshToken = errors.New("failed to create refresh token")
	ErrPasswordMismatch   = errors.New("password mismatch")
	ErrPasswordTooShort   = errors.New("password too short, minimum 8 characters")
	ErrPasswordTooLong    = errors.New("password too long, maximum 20 characters")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailNotFound      = errors.New("email not found")
	ErrPasswordWeak       = errors.New("password must contain at least one uppercase letter and one number")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenInvalid       = errors.New("token invalid")
)

type (
	RegisterRequest struct {
		Email    string `json:"email"    form:"email"    validate:"required,email"`
		Name     string `json:"name"     form:"name"     validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email"    form:"email"    validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}

	VerifyRequest struct {
		Token string `json:"token" form:"token"`
	}
)
