package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	model := new(models.Category)
	categories := new([]*models.Category)

	result := db.Ctx.Find(categories, model)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(categories)
}
