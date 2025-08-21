package server

import (
	"errors"
	"microservice/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get Service
// @Tags service
// @Accept json
// @Produce json
// @Param name query string false "Service name"
// @Success 200 {array} models.Service "response"
// @Router /api/services [get]
func (s *Server) GetServices(c *fiber.Ctx) error {
	name := c.Query("name")
	resp, err := s.Deps.PG.GetServices(name)
	if err != nil {
		return s.InternalServerError(c, err)
	}
	return c.JSON(resp)
}

// @Summary Create Service
// @Tags service
// @Accept json
// @Produce json
// @Param data body models.CreateServicesRequest true "Service creation payload"
// @Success 200 {object} models.Service "response"
// @Router /api/services/create [post]
func (s *Server) CreateServices(c *fiber.Ctx) error {
	var req models.CreateServicesRequest

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

// @Summary Update Service
// @Tags service
// @Accept json
// @Produce json
// @Param data body models.UpdateServicesRequest true "Service update payload"
// @Success 200 {object} models.Service "response"
// @Router /api/services/update [put]
func (s *Server) UpdateServices(c *fiber.Ctx) error {
	var req models.UpdateServicesRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	service := models.Service(req)

	if err := s.Deps.PG.UpdateServices(service); err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(service)
}

// @Summary Delete Service
// @Tags service
// @Accept json
// @Produce json
// @Param service_id query string true "Service ID"
// @Success 200 "OK"
// @Router /api/services/delete [delete]
func (s *Server) DeleteServices(c *fiber.Ctx) error {
	serviceID := c.Query("service_id")
	if serviceID == "" {
		return s.BadRequest(c, errors.New("service_id is required"))
	}

	err := s.Deps.PG.DeleteServices(serviceID)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
