package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetSubproducts(c *fiber.Ctx) error {
	data, err := s.Deps.Telebon.GetSubproducts()
	if err != nil {
		return s.InternalServerError(c, err)
	}

	if b, ok := data.([]byte); ok {
		return c.Send(b)
	}

	return c.JSON(data)
}
