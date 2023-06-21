package book

import (
	"sufrimiento/app/modules/book/handlers"

	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	app := fiber.New()

	app.Get("/books", handlers.GetBooks)

	app.Get("/book/:id", handlers.GetBook)

	app.Post("/book", handlers.PostBook)

	app.Put("/book/:id", handlers.PutBook)

	app.Delete("/book/:id", handlers.DeleteBook)

	return app
}
