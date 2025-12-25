package handlers

import (
	"procurement-system/config"
	"procurement-system/models"
	"procurement-system/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	req.Password = string(hash)

	config.DB.Create(&req)
	return c.JSON(fiber.Map{"message": "Register success"})
}

func Login(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if result := config.DB.Where("username = ?", req.Username).First(&user); result.Error != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Wrong password"})
	}

	token, _ := utils.GenerateToken(user.ID, user.Role)
	return c.JSON(fiber.Map{"token": token})
}
