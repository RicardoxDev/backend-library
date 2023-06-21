package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostAuthor(c *fiber.Ctx) error {
	a := new(models.Author)

	if err := c.BodyParser(a); err != nil {
		return err
	}

	result := db.Ctx.Create(a)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(201).JSON(a)
}
