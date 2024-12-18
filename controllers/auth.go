package controllers

import (
	"context"
	"event_management/database"
	"event_management/models"
	"event_management/utils"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error{
	var user models.User

	if err:= c.BodyParser(&user); err!=nil{
        log.Println("check 1")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"Invalid request"})
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost) 
	if err!=nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to hash password"})
	} 
	user.Password=string(hashedPassword)

	collection := database.DB.Collection("users")
	_,err = collection.InsertOne(context.Background(),user)

	if err!=nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":"Failed to register user"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})

}

func Login(c *fiber.Ctx)error{
	var credentials models.User
    if err := c.BodyParser(&credentials); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Find the user by email
    collection := database.DB.Collection("users")
    var user models.User
    err := collection.FindOne(context.Background(), bson.M{"email": credentials.Email}).Decode(&user)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Compare passwords
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.Email, user.Role)
    if err != nil {
		log.Println(user.Email)
		log.Println(user.Role)
		log.Println(err)
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Login successful", "token": token})

}