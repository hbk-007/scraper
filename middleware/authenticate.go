package middleware

import "github.com/gofiber/fiber/v2"

const TOKEN = "your_static_token"

func Authenticate(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token != TOKEN {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid or missing token"})
	}
	return c.Next()
}
