package middleware

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTProtected is a middleware function that checks for a valid JWT token in the Authorization header.
// If the token is valid, it allows the request to proceed; otherwise, it returns an unauthorized error.
func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}
		// Extract the token from the Authorization header
		tokenStr := authHeader[len("Bearer "):]

		// Parse the token using the secret key from the environment variable
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})// Parse the token using the secret key from the environment variable

        // Check if the token is valid and not expired
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		return c.Next()// Proceed to the next handler if the token is valid
	}
}
