package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/fiber-url_shortener/controller"
)

func SetupRoutes(app *fiber.App) {
	
	app.Post("/shorten", controller.UrlShorten)
}