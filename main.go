package main

import (
	"event_management/database"
	"event_management/routes"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in loading the env file")
	}

	database.ConnectDB()

	app:= fiber.New()

	routes.SetUpRoutes(app)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))

}