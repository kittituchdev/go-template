package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittituchdev/go-template/pkg"
)

func main() {
	app := fiber.New()

	pkg.SetupRoutes(app)

	app.Listen(":3000")
}
