package server

import (
	"microservice/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get Program
// @Tags program
// @Accept json
// @Produce json
// @Param type query string false "Program type"
// @Param name query string false "Program name"
// @Param active query boolean false "Active flag"
// @Success 200 {array} models.Program "response"
// @Router /api/program [get]
func (s *Server) GetProgram(c *fiber.Ctx) error {
	programType := c.Query("type")
	name := c.Query("name")
	active := c.QueryBool("active")

	resp, err := s.Deps.PG.GetProgram(programType, name, active)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

// @Summary Get Program
// @Tags program
// @Accept json
// @Produce json
// @Param id path string true "Program ID (UUID)"
// @Success 200 "OK"
// @Router /api/program/{id} [get]
func (s *Server) GetProgramByID(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := s.Deps.PG.GetProgramBYID(id)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

// @Summary Create Program
// @Tags program
// @Accept json
// @Produce json
// @Param data body models.CreateProgramRequest true "Program data"
// @Success 200 {object} models.Program "response"
// @Router /api/program/create [post]
func (s *Server) CreateProgram(c *fiber.Ctx) error {
	var req models.CreateProgramRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	finalPrice, finalDiscount := calculatePriceAndDiscount(req.FixedPrice, req.TotalServicesCost, req.DiscountPercent)

	program := models.Program{
		Type:              req.Type,
		Name:              req.Name,
		Image:             req.Image,
		FixedPrice:        finalPrice,
		TotalServicesCost: req.TotalServicesCost,
		DiscountPercent:   finalDiscount,
		ValidUntil:        req.ValidUntil,
		Terms:             req.Terms,
		Active:            req.Active,
	}

	resp, err := s.Deps.PG.CreateProgram(program)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

// @Summary Update Program
// @Tags program
// @Accept json
// @Produce json
// @Param data body models.UpdateProgramRequest true "Program data"
// @Success 200 {object} models.Program "response"
// @Router /api/program/update [put]
func (s *Server) UpdateProgram(c *fiber.Ctx) error {
	var req models.UpdateProgramRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	finalPrice, finalDiscount := calculatePriceAndDiscount(req.FixedPrice, req.TotalServicesCost, req.DiscountPercent)

	program := models.Program{
		ID:                req.ID,
		Type:              req.Type,
		Name:              req.Name,
		Image:             req.Image,
		FixedPrice:        finalPrice,
		TotalServicesCost: req.TotalServicesCost,
		DiscountPercent:   finalDiscount,
		ValidUntil:        req.ValidUntil,
		Terms:             req.Terms,
		Active:            req.Active,
	}

	err := s.Deps.PG.UpdateProgram(program)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(program)
}

// @Summary Delete Program
// @Tags program
// @Accept json
// @Produce json
// @Param id path string true "Program ID (UUID)"
// @Success 200 "OK"
// @Router /api/program/{id} [delete]
func (s *Server) DeleteProgram(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := s.Deps.PG.DeleteProgram(id); err != nil {
		return s.InternalServerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
