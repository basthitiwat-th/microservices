package main

import (
	"log"
	"microservices/config"
	handler "microservices/handlers"
	"microservices/repositories"
	"microservices/services"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Init()
	config.InitTimeZone()
	db := config.InitDatabase()

	// Create Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Authen
	authRepositoryDB := repositories.NewUserRepository(db)
	authService := services.NewAuthenicationServices(authRepositoryDB)
	authHandler := handler.NewAuthHandler(authService)
	authRoute := app.Group("/api/auth")

	authRoute.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test auth")
	})
	authRoute.Post("/register", authHandler.RegisterUser)
	authRoute.Post("/login", authHandler.LoginUser)

	// Product
	productRepositoryDB := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepositoryDB)
	productHandler := handler.NewProductHandler(productService)
	productRoute := app.Group("/api/products")

	productRoute.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test product")
	})
	productRoute.Post("/create", productHandler.CreateProduct)
	productRoute.Get("/", productHandler.GetAllProducts)
	productRoute.Get("/:product_id", productHandler.GetProductById)

	// Order
	orderRepositoryDB := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepositoryDB, productRepositoryDB)
	orderHandler := handler.NewOrderHandler(orderService)
	orderRoute := app.Group("/api/orders", authHandler.AuthMiddleware)

	productRoute.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test order")
	})
	orderRoute.Post("/create", orderHandler.CreateOrder)
	orderRoute.Put("/cancel/:order_id", orderHandler.CancelOrder)
	orderRoute.Get("/detail/:order_id", orderHandler.GetOrderByOrderID)

	// User
	userRepositoryDB := repositories.NewUserRepository(db)
	userService := services.NewUserServices(userRepositoryDB, orderRepositoryDB)
	userHandler := handler.NewUserHandler(userService)
	userRoute := app.Group("/api/user", authHandler.AuthMiddleware)

	userRoute.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test order")
	})
	userRoute.Get("/profile", userHandler.GetProfile)
	userRoute.Get("/orderhistory", userHandler.GetOrderHistory)

	// Start server
	port := os.Getenv("SERVER_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
