package user

import "github.com/gofiber/fiber/v2"

type userHandler struct {
	userService Service
}

func NewHandler(userService Service) *userHandler {
	return &userHandler{userService}
}
func (s *userHandler) Register(c *fiber.Ctx) error {
	var input UserInput
	err := c.BodyParser(&input)
	if err != nil {
		c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
		return err
	}
	userData, err := s.userService.RegisterUser(input)
	if err != nil {
		c.JSON(fiber.Map{"status": "error", "message": err.Error(), "data": nil})
		return err
	}
	c.JSON(fiber.Map{"status": "success", "message": "user created", "data": userData})
	return nil
}

// func (s *userHandler) Login(c *fiber.Ctx) {

// }
