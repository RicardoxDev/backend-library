package handlers

import (
	"sufrimiento/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetComment(c *fiber.Ctx) error {
	comment, statusCode := services.GetCommentParsed(c.Params("id"))

	if comment == nil {
		return c.SendStatus(statusCode)
	}
	return c.Status(statusCode).JSON(comment)
}
