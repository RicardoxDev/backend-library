package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostComment(c *fiber.Ctx) error {
	comment := new(models.Comment)

	if err := c.BodyParser(comment); err != nil {
		return c.SendStatus(500)
	}

	result := db.Ctx.Create(comment)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(201).JSON(comment)
}
