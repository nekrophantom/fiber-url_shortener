package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/fiber-url_shortener/config"
	"github.com/nekrophantom/fiber-url_shortener/routes"
)

func welcome (c *fiber.Ctx) error {
	return c.JSON("Welcome to the url shortener API")
}

func main() {
	config.LoadConfig()

	app := fiber.New()

	app.Get("/", welcome)

	routes.SetupRoutes(app)

	err := app.Listen(":3001")
	
	if err != nil {
		panic(err)
	}
}