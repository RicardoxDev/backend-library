package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteAuthor(c *fiber.Ctx) error {
	result := db.Ctx.Delete(&models.Author{}, c.Params("id"))

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
