package models

type Author struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Books []Book `gorm:"many2many:author_books;" json:"books"`
}
