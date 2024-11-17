package routes

import (
    "github.com/gofiber/fiber/v2"
    "event_management/controllers"
    "event_management/middleware"
)

func SetUpRoutes(app *fiber.App) {

    app.Post("/register", controllers.Register)
    app.Post("/login", controllers.Login)

    app.Get("/events", middleware.AuthMiddleware, controllers.GetEvents)
    app.Post("/events", middleware.AuthMiddleware, controllers.CreateEvent)
}
