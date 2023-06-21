package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PutComment(c *fiber.Ctx) error {
	dbComment := new(models.Comment)
	comment := new(models.Comment)
	id := c.Params("id")

	if err := c.BodyParser(comment); err != nil {
		return err
	}

	result := db.Ctx.First(dbComment, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	result = db.Ctx.Model(dbComment).Updates(comment)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(202)
}
