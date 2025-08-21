package server

import (
	"microservice/internal/deps"
	"microservice/internal/middleware"

	_ "microservice/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	swagger "github.com/swaggo/fiber-swagger"
)

type Server struct {
	App  *fiber.App
	Deps *deps.Deps
}

func New(deps *deps.Deps) *Server {
	s := &Server{
		App: fiber.New(fiber.Config{}),

		Deps: deps,
	}

	s.App.Use(cors.New())

	s.App.Get("/healthz", s.healthzHandler)

	// "microservice/internal/middleware"
	api := s.App.Group("/api", middleware.Middleware())
	//api := s.App.Group("/api") // localhost:8080/api

	s.App.Get("/swagger/*", swagger.WrapHandler)

	// products
	api.Get("/subproducts", s.GetSubproducts)

	// programs
	api.Get("/program", s.GetProgram)
	api.Get("/program/:id", s.GetProgramByID)
	api.Post("/program/create", s.CreateProgram)
	api.Put("/program/update", s.UpdateProgram)
	api.Delete("/program/:id", s.DeleteProgram)

	// services
	api.Get("/services", s.GetServices)
	api.Post("/services/create", s.CreateServices)
	api.Put("/services/update", s.UpdateServices)
	api.Delete("/services/:id", s.DeleteServices)

	// program services
	api.Get("/program_services", s.GetProgramServices)
	api.Post("/program_services/create", s.CreateProgramService)
	api.Delete("/program_services/:id", s.DeleteProgramService)

	return s
}

func (s *Server) healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

type responce struct {
	Message string `json:"message"`
}

func (s *Server) ResponceOK(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(responce{"ok"})
}

func (s *Server) InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).
		JSON(responce{err.Error()})
}

func (s *Server) BadRequest(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(responce{err.Error()})
}

func (s *Server) Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(responce{fiber.ErrUnauthorized.Message})
}
