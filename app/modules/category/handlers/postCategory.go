package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostCategory(c *fiber.Ctx) error {
	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return err
	}

	result := db.Ctx.Create(category)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}
