package server

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Get Subproducts
// @Tags Subproducts
// @Accept json
// @Produce json
// @Success 200 {object} interface{} "response from Telebon"
// @Router /api/subproducts [get]
func (s *Server) GetSubproducts(c *fiber.Ctx) error {
	data, err := s.Deps.Telebon.GetSubproducts()
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(data)
}
