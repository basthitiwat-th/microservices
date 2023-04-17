package handler

import (
	"microservices/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

type UserHandler interface {
	GetProfile(c *fiber.Ctx) error
	GetOrderHistory(c *fiber.Ctx) error
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

// Handler ดึง Profile ของ User
func (h *userHandler) GetProfile(c *fiber.Ctx) error {
	userid := c.Locals("user_id").(uint)

	userProfile, err := h.userService.GetProfile(userid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to get Profile"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"profile": userProfile})
}

// Handler ดึงรายละเอียดของ order history
func (h *userHandler) GetOrderHistory(c *fiber.Ctx) error {
	userid := c.Locals("user_id").(uint)

	orderhistory, err := h.userService.GetOrderHistoryByUserID(uint(userid))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to get OrderHistory"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"orders": orderhistory})
}
