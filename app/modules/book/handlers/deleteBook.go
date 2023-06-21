package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteBook(c *fiber.Ctx) error {
	result := db.Ctx.Delete(&models.Book{}, c.Params("id"))

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
