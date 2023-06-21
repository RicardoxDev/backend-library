package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PutUser(c *fiber.Ctx) error {
	u := new(models.User)
	dbUser := new(models.User)
	id := c.Params("id")

	if err := c.BodyParser(u); err != nil {
		return err
	}

	result := db.Ctx.First(&dbUser, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	result = db.Ctx.Model(&dbUser).Updates(u)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
