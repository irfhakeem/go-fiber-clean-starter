package dto

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/constants"
)

type (
	JWTPayload struct {
		UserId int64              `json:"user_id"`
		Role   constants.UserRole `json:"role"`
	}

	JWTClaims struct {
		UserId int64              `json:"user_id"`
		Role   constants.UserRole `json:"role"`
		jwt.RegisteredClaims
	}
)
