package models

type Category struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Tag   string `json:"tag"`
	Books []Book `gorm:"many2many:category_books;" json:"books"`
}
