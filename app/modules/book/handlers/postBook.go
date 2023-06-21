package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostBook(c *fiber.Ctx) error {
	b := new(models.Book)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	result := db.Ctx.Create(b)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(201).JSON(b)
}
