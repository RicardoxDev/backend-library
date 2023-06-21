package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PutAuthor(c *fiber.Ctx) error {
	a := new(models.Author)
	dbAuthor := new(models.Author)
	id := c.Params("id")

	if err := c.BodyParser(a); err != nil {
		return c.SendStatus(500)
	}

	result := db.Ctx.First(dbAuthor, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	result = db.Ctx.Model(dbAuthor).Updates(a)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.Status(202).JSON(dbAuthor)
}
