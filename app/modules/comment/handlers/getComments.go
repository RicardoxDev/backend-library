package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetComments(c *fiber.Ctx) error {
	model := new(models.Comment)
	comments := new([]*models.Comment)

	result := db.Ctx.Find(comments, model)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(comments)
}
