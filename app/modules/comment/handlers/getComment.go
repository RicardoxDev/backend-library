package handlers

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	id := c.Params("id")

	result := db.Ctx.First(comment, id)

	if result.Error != nil {
		return c.SendStatus(404)
	}

	cParsed := new(models.CommentDTO)
	comment.ParseToDTO(cParsed)

	b := new(models.Book)
	db.Ctx.First(b, comment.BookID)

	bParsed := new(models.BookDTO)
	b.ParseToDTO(bParsed)

	u := new(models.User)
	db.Ctx.First(u, comment.UserID)

	uParsed := new(models.UserDTO)
	u.ParseToDTO(uParsed)

	cParsed.Book = *bParsed
	cParsed.User = *uParsed
	return c.Status(200).JSON(cParsed)
}
