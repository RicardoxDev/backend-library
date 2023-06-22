package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"
	"sufrimiento/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetComments(c *fiber.Ctx) error {
	model := new(models.Comment)
	comments := []*models.Comment{}

	result := db.Ctx.Find(&comments, model)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	commentsParsed := make([]*models.CommentDTO, len(comments))

	for i := range commentsParsed {
		commentParsed, statusCode := services.GetCommentParsed(comments[i].ID)

		if commentParsed == nil {
			return c.SendStatus(statusCode)
		}

		commentsParsed[i] = commentParsed
	}

	return c.Status(200).JSON(commentsParsed)
}
