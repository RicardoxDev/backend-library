package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	u := c.Params("id")

	result := db.Ctx.Delete(&models.User{}, u)

	if result.Error != nil {
		c.SendStatus(404)
	}

	return c.SendStatus(200)
}
