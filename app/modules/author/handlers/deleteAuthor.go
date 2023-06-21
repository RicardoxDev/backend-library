package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteAuthor(c *fiber.Ctx) error {
	a := new(models.Author)
	id := c.Params("id")

	result := db.Ctx.Delete(a, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
