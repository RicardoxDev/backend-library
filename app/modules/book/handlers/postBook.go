package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return err
	}

	result := db.Ctx.Create(book)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}
