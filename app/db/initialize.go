package db

import (
	"log"
	"sufrimiento/app/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Ctx *gorm.DB

func StartDB() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{}, &models.Category{}, &models.Comment{})

	if err != nil {
		log.Fatal(err)
	}

	Ctx = db
}
