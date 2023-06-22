package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"
	"sufrimiento/app/services"

	"github.com/gofiber/fiber/v2"
)

func PostUser(c *fiber.Ctx) error {
	userDTO := new(models.CreateUserDTO)

	if err := c.BodyParser(userDTO); err != nil {
		return err
	}

	user := models.User{Name: userDTO.Name, Password: userDTO.Password}

	result := db.Ctx.Create(&user)

	if result.Error != nil {
		c.SendStatus(500)
	}

	userParsed, statusCode := services.GetUserParsed(user.ID)

	if userParsed == nil {
		return c.SendStatus(statusCode)
	}

	return c.Status(statusCode).JSON(userParsed)
}
