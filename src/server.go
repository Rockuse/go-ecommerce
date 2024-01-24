package server

import (
	"github.com/Rockuse/go-ecommerce/src/app/core/routes"
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
	api := router.Group("/api")
	v1 := api.Group("/v1")
	db := s.db
	for _, m := range routes.RegisteredModules() {
		m.Routes(&v1, db)
	} 

}
