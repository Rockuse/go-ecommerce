package routes

import (
	"github.com/gofiber/fiber/v2"
)

var UserModule = Module{
	Name: "user",
	Routes: func(api *fiber.Group) {
		userApi := api.Group("/user")

		userApi.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("test")
		})
	},
}
