package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/go-template/internal/example"
)

func SetupRoutes(app *fiber.App) {
	// group for example
	exampleGroup := app.Group("/example")
	{
		exampleGroup.Post("/", example.CreateExample)
		exampleGroup.Get("/", example.GetAllExample)
		exampleGroup.Get("/", example.GetExample)
		exampleGroup.Put("/", example.UpdateExample)
		exampleGroup.Delete("/", example.DeleteExample)
	}

}
