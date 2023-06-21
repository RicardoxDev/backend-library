package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	id := c.Params("id")

	result := db.Ctx.First(category, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(category)
}
