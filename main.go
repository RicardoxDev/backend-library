package main

import (
	"sufrimiento/app/db"
	"sufrimiento/app/modules/author"
	"sufrimiento/app/modules/book"
	"sufrimiento/app/modules/category"
	"sufrimiento/app/modules/comment"
	"sufrimiento/app/modules/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Test Backend Library",
	})

	db.StartDB()

	app.Use(logger.New())

	UsersApp := user.Controller()
	app.Mount("/", UsersApp)

	AuthorsApp := author.Controller()
	app.Mount("/", AuthorsApp)

	BooksApp := book.Controller()
	app.Mount("/", BooksApp)

	CategoriesApp := category.Controller()
	app.Mount("/", CategoriesApp)

	CommentsApp := comment.Controller()
	app.Mount("/", CommentsApp)

	app.Listen(":8080")
}
