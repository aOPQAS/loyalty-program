package server

import (
	"microservice/pkg/models"

	"github.com/gofiber/fiber/v2"
)

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

func (s *Server) GetProgramByID(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := s.Deps.PG.GetProgramBYID(id)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

type createProgramRequest struct {
	Type              string `json:"type"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	FixedPrice        int    `json:"fixed_price"`
	TotalServicesCost int    `json:"total_services_cost"`
	DiscountPercent   int    `json:"discount_percent"`
	ValidUntil        string `json:"valid_until"`
	Terms             string `json:"terms"`
	Active            bool   `json:"active"`
}

func (s *Server) CreateProgram(c *fiber.Ctx) error {
	var req createProgramRequest

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

type updateProgramRequest struct {
	ID                string `json:"id"`
	Type              string `json:"type"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	FixedPrice        int    `json:"fixed_price"`
	TotalServicesCost int    `json:"total_services_cost"`
	DiscountPercent   int    `json:"discount_percent"`
	ValidUntil        string `json:"valid_until"`
	Terms             string `json:"terms"`
	Active            bool   `json:"active"`
}

func (s *Server) UpdateProgram(c *fiber.Ctx) error {
	var req updateProgramRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	finalPrice, finalDiscount := calculatePriceAndDiscount(req.FixedPrice, req.TotalServicesCost, req.DiscountPercent)

	err := s.Deps.PG.UpdateProgram(models.Program{
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
	})
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Server) DeleteProgram(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := s.Deps.PG.DeleteProgram(id); err != nil {
		return s.InternalServerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
