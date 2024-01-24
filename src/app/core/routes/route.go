package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Module struct {
	Name   string
	Routes func(*fiber.Router, *gorm.DB)
}

func RegisteredModules() []Module {
	return []Module{UserModule}
}
