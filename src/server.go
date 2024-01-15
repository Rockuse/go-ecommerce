package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	engine *fiber.App
	db     *gorm.DB
}

func NewServer(dbConnection *gorm.DB) *Server {
	return &Server{
		engine: fiber.New(),
		db:     dbConnection,
	}
}

func (s *Server) Run() error {
	s.RoutesInit()
	return s.engine.Listen(":3000")
}

func (s *Server) RoutesInit() {
	router := s.engine
	api := router.Group("/api/v1")
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
