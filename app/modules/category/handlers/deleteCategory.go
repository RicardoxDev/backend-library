package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(c *fiber.Ctx) error {
	result := db.Ctx.Delete(&models.Category{}, c.Params("id"))

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
