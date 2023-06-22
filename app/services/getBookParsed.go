package services

import (
	"sufrimiento/app/db"
	"sufrimiento/app/models"
)

func GetBookParsed(id interface{}) (*models.BookDTO, int) {
	book := new(models.Book)
	result := db.Ctx.First(book, id)

	if result.Error != nil {
		return nil, 404
	}

	bookParsed := new(models.BookDTO)
	book.ParseToDTO(bookParsed)

	return bookParsed, 200
}
