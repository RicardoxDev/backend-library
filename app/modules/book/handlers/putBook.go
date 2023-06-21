package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PutBook(c *fiber.Ctx) error {
	dbBook := new(models.Book)
	book := new(models.Book)
	id := c.Params("id")

	if err := c.BodyParser(book); err != nil {
		return err
	}

	result := db.Ctx.First(dbBook, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	result = db.Ctx.Model(dbBook).Updates(book)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(dbBook)
}
