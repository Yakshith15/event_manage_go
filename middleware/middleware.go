package middleware

import (
	"event_management/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error{
	token := c.Get("Authorization")
	if token == ""{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"No token provided"})
	}

	claims,err := utils.ParseJWT(strings.TrimPrefix(token,"Bearer "))
	if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid or expired token"})
	}

	c.Locals("user",claims)

	return c.Next()
}