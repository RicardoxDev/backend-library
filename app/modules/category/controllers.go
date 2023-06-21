package category

import (
	"sufrimiento/app/modules/category/handlers"

	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	app := fiber.New()

	app.Get("/categories", handlers.GetCategories)

	app.Get("/category/:id", handlers.GetCategory)

	app.Post("/category", handlers.PostCategory)

	app.Put("/category/:id", handlers.PutCategory)

	app.Delete("/category/:id", handlers.DeleteCategory)

	return app
}
