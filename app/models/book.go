package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title      string     `json:"title"`
	Authors    []Author   `gorm:"many2many:author_books;" json:"authors"`
	Categories []Category `gorm:"many2many:category_books;" json:"categories"`
	Comments   []Comment  `json:"comments"`
}

type BookDTO struct {
	ID      uint
	Title   string
	Authors []Author
}

func (book *Book) ParseToDTO(dto *BookDTO) error {
	dto.ID = book.ID
	dto.Title = book.Title
	dto.Authors = book.Authors

	return nil
}
