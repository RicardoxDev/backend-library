package comment

import (
	"sufrimiento/app/modules/comment/handlers"

	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	app := fiber.New()

	app.Get("/comments", handlers.GetComments)

	app.Get("/comment/:id", handlers.GetComment)

	app.Post("/comment", handlers.PostComment)

	app.Put("/comment/:id", handlers.PutComment)

	app.Delete("/comment/:id", handlers.DeleteComment)

	return app
}
