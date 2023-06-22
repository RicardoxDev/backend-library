package services

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"
)

func GetCommentParsed(id interface{}) (*models.CommentDTO, int) {
	comment := new(models.Comment)
	dto := new(models.CommentDTO)

	result := db.Ctx.First(comment, id)

	if result.Error != nil {
		return nil, 404
	}

	dto.ID = comment.ID

	userParsed, userStatusCode := GetUserParsed(comment.UserID)

	if userParsed == nil {
		return nil, userStatusCode
	}

	dto.User = *userParsed
	dto.Body = comment.Body

	bookParsed, bookStatusCode := GetBookParsed(comment.BookID)

	if bookParsed == nil {
		return nil, bookStatusCode
	}

	dto.Book = *bookParsed

	return dto, 200
}
