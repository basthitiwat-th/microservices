package handler

import (
	"microservices/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderService services.OrderService
}

type OrderHandler interface {
	CreateOrder(c *fiber.Ctx) error
	CancelOrder(c *fiber.Ctx) error
	GetOrderByOrderID(c *fiber.Ctx) error
}

func NewOrderHandler(orderService services.OrderService) OrderHandler {
	return &orderHandler{orderService: orderService}
}

// Handler สร้าง Order
func (h *orderHandler) CreateOrder(c *fiber.Ctx) error {
	var orderReq services.OrderRequest
	if err := c.BodyParser(&orderReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	if err := h.orderService.CreateOrder(orderReq); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to create Order"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Order created successfully"})
}

// Handler ยกเลิก Order ตาม Id
func (h *orderHandler) CancelOrder(c *fiber.Ctx) error {
	orderID, err := c.ParamsInt("order_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	err = h.orderService.CancelOrder(uint(orderID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to Cancel Order"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "CancelOrder successfully"})

}

// Handeler ดึงข้อมูล order ตาม ID
func (h *orderHandler) GetOrderByOrderID(c *fiber.Ctx) error {
	orderID, err := c.ParamsInt("order_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	order, err := h.orderService.GetOrderByOrderID(uint(orderID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to Cancel Order"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"order": order})
}
