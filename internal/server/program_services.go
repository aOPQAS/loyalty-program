package server

import (
	"microservice/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetProgramServices(c *fiber.Ctx) error {
	programID := c.Query("program_id")

	resp, err := s.Deps.PG.GetProgramServices(programID)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return c.JSON(resp)
}

type createProgramServiceRequest struct {
	ProgramID string `json:"program_id"`
	ServiceID string `json:"service_id"`
}

func (s *Server) CreateProgramService(c *fiber.Ctx) error {
	var req createProgramServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	if err := s.Deps.PG.CreateProgramService(models.ProgramService{
		ProgramID: req.ProgramID,
		ServiceID: req.ServiceID,
	}); err != nil {
		return s.InternalServerError(c, err)
	}

	return s.ResponceOK(c)
}

type deleteProgramServiceRequest struct {
	ProgramID string `json:"program_id"`
	ServiceID string `json:"service_id"`
}

func (s *Server) DeleteProgramService(c *fiber.Ctx) error {
	var req deleteProgramServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return s.BadRequest(c, err)
	}

	err := s.Deps.PG.DeleteProgramServices(req.ProgramID, req.ServiceID)
	if err != nil {
		return s.InternalServerError(c, err)
	}

	return s.ResponceOK(c)
}
