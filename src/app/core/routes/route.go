package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	Name   string
	Routes func(*fiber.Group)
}

func RegisteredModules() []Module {
	return []Module{}
}
