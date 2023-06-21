package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	books := new([]*models.Book)
	model := new(models.Book)

	result := db.Ctx.Find(books, model)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(books)
}
