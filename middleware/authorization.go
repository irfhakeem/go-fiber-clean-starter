package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/constants"
	"github.com/irfhakeem/go-fiber-clean-starter/helpers/utils"
	"github.com/irfhakeem/go-fiber-clean-starter/service"
)

func Authorize(jwtService service.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				utils.FailedResponse(dto.FAILED_UNAUTHORIZED, dto.FAILED_TOKEN_NOT_FOUND),
			)
		}

		if !strings.Contains(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(
				utils.FailedResponse(dto.FAILED_UNAUTHORIZED, dto.FAILED_TOKEN_INVALID),
			)
		}

		token := strings.Split(authHeader, "Bearer ")[1]
		payload, err := jwtService.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				utils.FailedResponse(dto.FAILED_UNAUTHORIZED, dto.FAILED_TOKEN_INVALID),
			)
		}

		claimsMap, ok := payload.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(
				utils.FailedResponse(dto.FAILED_UNAUTHORIZED, dto.FAILED_TOKEN_INVALID),
			)
		}

		claims := &dto.JWTClaims{
			UserId: int64(claimsMap["user_id"].(float64)),
			Role:   constants.UserRole(claimsMap["role"].(string)),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Unix(int64(claimsMap["exp"].(float64)), 0)),
			},
		}

		if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now()) {
			return c.Status(fiber.StatusUnauthorized).JSON(
				utils.FailedResponse(dto.FAILED_UNAUTHORIZED, dto.FAILED_TOKEN_EXPIRED),
			)
		}

		c.Locals("user_id", strconv.FormatInt(claims.UserId, 10))
		c.Locals("user_role", string(claims.Role))

		return c.Next()
	}
}
