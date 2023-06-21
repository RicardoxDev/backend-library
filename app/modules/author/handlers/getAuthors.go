package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetAuthors(c *fiber.Ctx) error {
	model := new(models.Author)
	authors := new([]*models.Author)

	result := db.Ctx.Find(authors, model)

	if result.Error != nil {
		return c.SendStatus(204)
	}

	return c.Status(200).JSON(authors)
}
