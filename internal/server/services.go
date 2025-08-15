package server

import (
	"microservice/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetServices(c *fiber.Ctx) error {
	name := c.Query("name")
	resp, err := s.Deps.PG.GetServices(name)
	if err != nil {
		return s.InternalServerError(c, err)
	}
	return c.JSON(resp)
}

type createServicesRequest struct {
	Name     string `json:"name"`
	Tarif    int    `json:"tarif"`
	Duration int    `json:"duration"`
}

func (s *Server) CreateServices(c *fiber.Ctx) error {
	var req createServicesRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	resp, err := s.Deps.PG.CreateServices(models.Service{
		Name:     req.Name,
		Tarif:    req.Tarif,
		Duration: req.Duration,
	})
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

type updateServicesRequest struct {
	ServiceID string `json:"service_id"`
	Name      string `json:"name"`
	Tarif     int    `json:"tarif"`
	Duration  int    `json:"duration"`
}

func (s *Server) UpdateServices(c *fiber.Ctx) error {
	var req updateServicesRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	err := s.Deps.PG.UpdateServices(models.Service{
		ServiceID: req.ServiceID,
		Name:      req.Name,
		Tarif:     req.Tarif,
		Duration:  req.Duration,
	})
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return s.ResponceOK(c)
}

type deleteServicesRequest struct {
	ServiceID string `json:"service_id"`
}

func (s *Server) DeleteServices(c *fiber.Ctx) error {
	var req deleteServicesRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	err := s.Deps.PG.DeleteServices(req.ServiceID)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return s.ResponceOK(c)
}
