package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PutCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	dbCategory := new(models.Category)
	id := c.Params("id")

	if err := c.BodyParser(category); err != nil {
		return c.SendStatus(500)
	}

	result := db.Ctx.First(dbCategory, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	result = db.Ctx.Model(dbCategory).Updates(category)

	if result.Error != nil {
		return c.SendStatus(408)
	}

	return c.SendStatus(200)
}
