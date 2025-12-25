package routes

import (
	"procurement-system/handlers"
	"procurement-system/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	api := app.Group("/api", middleware.Protected())
	api.Get("/items", handlers.GetItems)
	api.Post("/items", handlers.CreateItem)
	api.Get("/suppliers", handlers.GetSuppliers)
	api.Post("/suppliers", handlers.CreateSupplier)
	api.Post("/purchasings", handlers.CreatePurchase)
}
