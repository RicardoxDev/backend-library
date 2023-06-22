package handlers

import (
	"sufrimiento/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	user, statusCode := services.GetUserParsed(c.Params("id"))

	if user == nil {
		return c.SendStatus(statusCode)
	}

	return c.Status(statusCode).JSON(user)
}
