package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string     `json:"name"`
	Password           string     `json:"password"`
	FavoriteAuthors    []Author   `gorm:"many2many:favorite_authors;" json:"favoriteauthors"`
	FavoriteBooks      []Book     `gorm:"many2many:favorite_books;" json:"favoritebooks"`
	FavoriteCategories []Category `gorm:"many2many:favorite_categories;" json:"favoritecategories"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID                 uint
	Name               string
	FavoriteAuthors    []Author
	FavoriteBooks      []Book
	FavoriteCategories []Category
}

func (user *User) ParseToDTO(dto *UserDTO) error {
	dto.ID = user.ID
	dto.Name = user.Name
	dto.FavoriteAuthors = user.FavoriteAuthors
	dto.FavoriteBooks = user.FavoriteBooks
	dto.FavoriteCategories = user.FavoriteCategories

	return nil
}
