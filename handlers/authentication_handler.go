package handler

import (
	jwtPkg "microservices/pkg/jwt"
	"microservices/services"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService services.AuthenticationService
}

type AuthHandler interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	AuthMiddleware(c *fiber.Ctx) error
}

func NewAuthHandler(authService services.AuthenticationService) AuthHandler {
	return &authHandler{authService: authService}

}

// Handler สำหรับสมัครสมาชิก
func (h *authHandler) RegisterUser(c *fiber.Ctx) error {
	var registerRequest services.UserRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	if registerRequest.Username == "" || registerRequest.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Username and password are required"})
	}

	if err := h.authService.Register(registerRequest); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Unable to Register"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// Handler สำหรับเข้าสู่ระบบ
func (h *authHandler) LoginUser(c *fiber.Ctx) error {
	var loginRequest services.UserLogin
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Username and password are required"})
	}

	token, err := h.authService.Login(loginRequest)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid username or password"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})

}

func (h *authHandler) AuthMiddleware(c *fiber.Ctx) error {
	// ดึงค่า token จาก header
	authHeader := c.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	//validate token
	claim, ok := jwtPkg.ValidAndGetClaims(tokenString)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Token"})
	}
	c.Locals("user_id", claim.Id)

	return c.Next()
}
