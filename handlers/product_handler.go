package handler

import (
	"microservices/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productServices services.ProductServices
}

type ProductHandler interface {
	CreateProduct(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
}

func NewProductHandler(productServices services.ProductServices) ProductHandler {
	return &productHandler{productServices: productServices}
}

// Handler สร้างสินค้า
func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	var productReq services.ProductRequest
	if err := c.BodyParser(&productReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	if err := h.productServices.CreateProduct(productReq); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to create Product"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Product created successfully"})
}

// Handler ดึงรายการสินค้าทั้งหมด
func (h *productHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.productServices.GetAllProducts()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Product not found"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"products": products})

}

// Handler ดึงรายการสินค้าตาม ID
func (h *productHandler) GetProductById(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("product_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}
	product, err := h.productServices.GetProductById(uint(productID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Product not found"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"product": product})

}
