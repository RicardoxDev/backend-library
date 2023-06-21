package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func PostUser(c *fiber.Ctx) error {
	userDTO := new(models.CreateUserDTO)

	if err := c.BodyParser(userDTO); err != nil {
		return err
	}

	u := models.User{Name: userDTO.Name, Password: userDTO.Password}

	result := db.Ctx.Create(&u)

	if result.Error != nil {
		c.SendStatus(500)
	}

	uParsed := new(models.UserDTO)
	u.ParseToDTO(uParsed)

	return c.Status(201).JSON(uParsed)
}
