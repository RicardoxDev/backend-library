package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"
	"sufrimiento/app/services"

	"github.com/gofiber/fiber/v2"
)

func PutUser(c *fiber.Ctx) error {
	user := new(models.User)
	dbUser := new(models.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(500)
	}

	result := db.Ctx.First(&dbUser, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	result = db.Ctx.Model(&dbUser).Updates(user)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	userParsed, statusCode := services.GetUserParsed(dbUser.ID)

	if userParsed == nil {
		return c.SendStatus(statusCode)
	}

	return c.Status(statusCode).JSON(userParsed)
}
