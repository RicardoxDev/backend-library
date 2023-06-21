package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteComment(c *fiber.Ctx) error {
	model := new(models.Comment)
	id := c.Params("id")

	result := db.Ctx.Delete(model, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
