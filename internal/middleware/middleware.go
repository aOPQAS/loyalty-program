package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		expectedToken := os.Getenv("ACCESS_TOKEN")
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != expectedToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		return c.Next()
	}
}
