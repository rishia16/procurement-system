package handlers

import (
	"procurement-system/config"
	"procurement-system/models"

	"github.com/gofiber/fiber/v2"
)

func GetSuppliers(c *fiber.Ctx) error {
	var s []models.Supplier
	config.DB.Find(&s)
	return c.JSON(s)
}

func CreateSupplier(c *fiber.Ctx) error {
	var s models.Supplier
	if err := c.BodyParser(&s); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	config.DB.Create(&s)
	return c.JSON(s)
}
