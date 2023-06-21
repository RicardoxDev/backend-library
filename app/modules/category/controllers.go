package category

import (
	"sufrimiento/app/models"

	"github.com/gofiber/fiber/v2"
)

func Controller() *fiber.App {
	app := fiber.New()

	app.Get("/categories", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/category", func(c *fiber.Ctx) error {
		return nil
	})

	app.Post("/category", func(c *fiber.Ctx) error {
		tag := new(models.Category)

		if err := c.BodyParser(tag); err != nil {
			return err
		}

		return c.JSON(tag)
	})

	app.Put("/category", func(c *fiber.Ctx) error {
		tag := new(models.Category)

		if err := c.BodyParser(tag); err != nil {
			return err
		}

		return c.JSON(tag)
	})

	app.Delete("/category", func(c *fiber.Ctx) error {
		tag := new(models.Category)

		if err := c.BodyParser(tag); err != nil {
			return err
		}

		return c.JSON(tag)
	})

	return app
}
