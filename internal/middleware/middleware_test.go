package middleware_test

import (
	"net/http/httptest"
	"os"
	"testing"

	"microservice/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	os.Setenv("ACCESS_TOKEN", "test-token")

	app := fiber.New()
	app.Use(middleware.Middleware())

	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{
			name:           "No Authorization header",
			authHeader:     "",
			expectedStatus: fiber.StatusUnauthorized,
		},
		{
			name:           "Wrong token",
			authHeader:     "Bearer wrong-token",
			expectedStatus: fiber.StatusUnauthorized,
		},
		{
			name:           "Correct token",
			authHeader:     "Bearer test-token",
			expectedStatus: fiber.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
