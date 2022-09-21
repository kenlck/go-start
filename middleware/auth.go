package middleware

import (
	"crm_backend/config"
	"crm_backend/response"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.SECRET),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.MissingJWTResponseBody())
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(response.InvalidJWTResponseBody())
}
