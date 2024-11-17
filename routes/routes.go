package routes

import "github.com/gofiber/fiber/v2"
import "event_management/controllers"

func SetUpRoutes(app *fiber.App) {
	app.Get("/events", controllers.GetEvents)
	app.Post("/events",controllers.CreateEvent)
}