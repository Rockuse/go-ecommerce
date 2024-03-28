package routes

import (
	"github.com/Rockuse/go-ecommerce/src/app/modules/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var UserModule = Module{
	Name: "user",
	Routes: func(Api *fiber.App, db *gorm.DB) {
		userRepo := user.NewRepository(db)
		userService := user.NewService(userRepo)
		userHandler := user.NewHandler(userService)
		userApi := Api.Group("api/v1/user")
		userApi.Post("/", userHandler.Register)
		// userApi.Get("/", userHandler.Login)
	},
}
