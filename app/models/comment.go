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
