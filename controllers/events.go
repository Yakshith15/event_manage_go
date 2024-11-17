package controllers

import (
	"context"
	"net/http"
	"time"

	"event_management/database"
	"event_management/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.DB.Collection("events").Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch events"})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx){
		var event models.Event
		if err:= cursor.Decode(&event); err!=nil{
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch events"})
		}
		events = append(events,event)
	}

	return c.JSON(events)
}


func CreateEvent(c *fiber.Ctx) error{
	var event models.Event 
	if err:= c.BodyParser(&event); err!=nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	
	_,err := database.DB.Collection("events").InsertOne(ctx,event);
	if(err!=nil){
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create event"})
	}

	return c.Status(http.StatusCreated).JSON(event)
	
}