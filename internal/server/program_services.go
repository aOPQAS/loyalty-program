package server

import (
	"errors"
	"microservice/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get ProgramService
// @Tags ProgramService
// @Accept json
// @Produce json
// @Param program_id query string true "Program ID (UUID)"
// @Success 200 {array} models.ProgramService "response"
// @Router /api/program_services [get]
func (s *Server) GetProgramServices(c *fiber.Ctx) error {
	programID := c.Query("program_id")

	resp, err := s.Deps.PG.GetProgramServices(programID)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

// @Summary Create ProgramService
// @Tags ProgramService
// @Accept json
// @Produce json
// @Param data body models.CreateProgramServiceRequest true "Program services data"
// @Success 200 {object} models.ProgramService "response"
// @Router /api/program_services/create [post]
func (s *Server) CreateProgramService(c *fiber.Ctx) error {
	var req models.CreateProgramServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	ps := models.ProgramService(req)

	if err := s.Deps.PG.CreateProgramService(ps); err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(ps)
}

// @Summary Delete ProgramService
// @Tags ProgramService
// @Accept json
// @Produce json
// @Param program_id query string true "Program ID"
// @Param service_id query string true "Service ID"
// @Success 200 "OK"
// @Router /api/program_services/delete [delete]
func (s *Server) DeleteProgramService(c *fiber.Ctx) error {
	programID := c.Query("program_id")
	serviceID := c.Query("service_id")

	if programID == "" || serviceID == "" {
		return s.BadRequest(c, errors.New("program_id and service_id are required"))
	}

	err := s.Deps.PG.DeleteProgramServices(programID, serviceID)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}
