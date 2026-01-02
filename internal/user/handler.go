package user

import "github.com/gofiber/fiber/v2"



type userHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	userResponses, err := h.userService.FindAllUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(userResponses)
}