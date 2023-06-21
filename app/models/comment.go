package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID uint   `json:"userid"`
	User   User   `json:"user"`
	Body   string `json:"body"`
	BookID uint   `json:"bookid"`
	Book   Book   `json:"book"`
}

type CommentDTO struct {
	ID   uint
	User UserDTO
	Body string
	Book BookDTO
}

func (comment *Comment) ParseToDTO(dto *CommentDTO) error {
	dto.ID = comment.ID

	uParsed := new(UserDTO)
	comment.User.ParseToDTO(uParsed)
	dto.User = *uParsed

	dto.Body = comment.Body

	bParsed := new(BookDTO)
	comment.Book.ParseToDTO(bParsed)
	dto.Book = *bParsed

	return nil
}
