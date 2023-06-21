package author

import (
	"sufrimiento/app/modules/author/handlers"

	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	app := fiber.New()

	app.Get("/authors", handlers.GetAuthors)

	app.Get("/author/:id", handlers.GetAuthor)

	app.Post("/author", handlers.PostAuthor)

	app.Put("/author/:id", handlers.PutAuthor)

	app.Delete("/author/:id", handlers.DeleteAuthor)

	return app
}
