package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	engine *fiber.App
}

func NewServer() *Server {
	return &Server{
		engine: fiber.New(),
	}
}

func (s *Server) Run() error {
	return s.engine.Listen(":3000")
}

func (s *Server) RoutesInit() {
	router := s.engine
	api := router.Group("/api/v1")
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
