package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostAuthor(c *fiber.Ctx) error {
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return err
	}

	result := db.Ctx.Create(author)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}
