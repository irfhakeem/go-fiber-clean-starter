package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/irfhakeem/go-fiber-clean-starter/dto"
)

type (
	JWTService interface {
		GenerateToken(payload dto.JWTPayload, exp time.Time) (string, error)
		ValidateToken(token string) (*jwt.Token, error)
	}

	jwtService struct {
		secret string
		issuer string
	}
)

func NewJwtService() JWTService {
	return &jwtService{
		secret: os.Getenv("JWT_SECRET"),
		issuer: "go-fiber-clean-starter",
	}
}

func (j *jwtService) GenerateToken(payload dto.JWTPayload, exp time.Time) (string, error) {
	claims := dto.JWTClaims{
		UserId: payload.UserId,
		Role:   payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError(
				"unexpected signing method",
				jwt.ValidationErrorSignatureInvalid,
			)
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}
