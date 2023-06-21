package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	result := db.Ctx.Delete(&models.User{}, c.Params("id"))

	if result.Error != nil {
		c.SendStatus(404)
	}

	return c.SendStatus(200)
}
