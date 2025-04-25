package main

import (
	"log"
	"github.com/ConDavison1/rise-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"  
	"github.com/ConDavison1/rise-api/handlers"
	"github.com/ConDavison1/rise-api/middleware"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()           
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	if err := db.Connect(); err != nil {
		log.Fatal("Cannot Connect to the database..", err)
	}

	app := fiber.New()
	// Initial route for testing the API's just to check if the server is running
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rise Mini Booking API")
	})

	
	app.Post("/login", handlers.Login)

	// Route to get bookings, protected by JWT middleware
	// This route will only be accessible if the user is authenticated with a valid JWT token
	app.Get("/bookings", middleware.JWTProtected(), handlers.GetBookings)

	// Route to create a booking, protected by JWT middleware
	app.Post("/bookings", middleware.JWTProtected(), handlers.CreateBooking)

	// Route to get all programs
	app.Get("/programs", handlers.GetPrograms)
	
	//get a program by id
	app.Get("/programs/:id", handlers.GetProgramByID)
	
	//app is listening on port 3000
	app.Listen(":3000")
}
