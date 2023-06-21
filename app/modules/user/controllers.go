package user

import (
	"sufrimiento/app/modules/user/handlers"

	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	app := fiber.New()

	app.Get("/users", handlers.GetUsers)

	app.Get("/user/:id", handlers.GetUser)

	app.Post("/user", handlers.PostUser)

	app.Put("/user/:id", handlers.PutUser)

	app.Delete("/user/:id", handlers.DeleteUser)

	return app
}
