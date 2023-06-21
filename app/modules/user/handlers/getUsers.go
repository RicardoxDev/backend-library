package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var u []*models.User

	result := db.Ctx.Find(&u, models.User{})

	if result.Error != nil {
		return c.SendStatus(500)
	}

	uParsed := make([]models.UserDTO, len(u))

	for i := range u {
		u[i].ParseToDTO(&uParsed[i])
	}

	return c.Status(200).JSON(uParsed)
}
