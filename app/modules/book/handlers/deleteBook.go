package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	model := new(models.Book)

	result := db.Ctx.Delete(model, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
