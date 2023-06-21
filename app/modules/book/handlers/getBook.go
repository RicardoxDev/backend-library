package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	b := new(models.Book)
	id := c.Params("id")

	result := db.Ctx.First(b, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(b)
}
