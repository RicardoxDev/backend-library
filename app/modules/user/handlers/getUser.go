package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	u := new(models.User)

	result := db.Ctx.First(&u, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	uDTO := new(models.UserDTO)
	u.ParseToDTO(uDTO)

	return c.Status(200).JSON(uDTO)
}
