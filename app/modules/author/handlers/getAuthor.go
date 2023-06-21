package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	a := new(models.Author)

	result := db.Ctx.First(&a, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(a)
}
