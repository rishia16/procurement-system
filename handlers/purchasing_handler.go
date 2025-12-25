package handlers

import (
	"bytes"
	"net/http"
	"time"

	"procurement-system/config"
	"procurement-system/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type PurchaseRequest struct {
	SupplierID uint `json:"supplier_id"`
	Items      []struct {
		ItemID uint `json:"item_id"`
		Qty    int  `json:"qty"`
	} `json:"items"`
}

func CreatePurchase(c *fiber.Ctx) error {
	userToken := c.Locals("user")
	if userToken == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id"})
	}
	userID := uint(userIDFloat)

	var req PurchaseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.SupplierID == 0 || len(req.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Supplier and items are required"})
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed start transaction"})
	}

	total := 0.0
	purchase := models.Purchasing{
		Date:       time.Now().Format("2006-01-02"),
		SupplierID: req.SupplierID,
		UserID:     userID,
	}

	if err := tx.Create(&purchase).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Failed create purchasing"})
	}

	for _, i := range req.Items {
		if i.Qty <= 0 {
			tx.Rollback()
			return c.Status(400).JSON(fiber.Map{"error": "Quantity must be > 0"})
		}

		var item models.Item
		if err := tx.First(&item, i.ItemID).Error; err != nil {
			tx.Rollback()
			return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
		}

		if item.Stock < i.Qty {
			tx.Rollback()
			return c.Status(400).JSON(fiber.Map{"error": "Stock not enough for item " + item.Name})
		}

		subTotal := float64(i.Qty) * item.Price
		total += subTotal

		detail := models.PurchasingDetail{
			PurchasingID: purchase.ID,
			ItemID:       item.ID,
			Qty:          i.Qty,
			SubTotal:     subTotal,
		}

		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{"error": "Failed create purchasing detail"})
		}

		item.Stock -= i.Qty
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{"error": "Failed update stock"})
		}
	}

	if err := tx.Model(&purchase).Update("grand_total", total).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Failed update grand total"})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Failed commit transaction"})
	}

	go http.Post(
		"https://webhook.site/xxxx",
		"application/json",
		bytes.NewBuffer([]byte(`{"status":"purchase_success"}`)),
	)

	return c.JSON(fiber.Map{
		"message":     "Purchase success",
		"grand_total": total,
	})
}
