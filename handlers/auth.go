package handlers

import (
	"log"
	"os"
	"time"

	"github.com/ConDavison1/rise-api/db"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)
// LoginReq is the request structure for login
// It contains the email and password fields that are required for authentication.
type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	var req LoginReq // Define the request structure for login
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})// Handle error if request body parsing fails
	}
	var dbPassword string
	err := db.DB.QueryRow(c.Context(), "SELECT password FROM users WHERE email = $1", req.Email).Scan(&dbPassword)// Check if the user exists in the database
	if err != nil || dbPassword != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ 
			"error": "Invalid Credentials",
		})// Handle error if user not found or password doesn't match
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))) // Use the JWT secret from the environment variable
	if err != nil {
		log.Fatal(err)
		// Handle error if token generation fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString})

}

