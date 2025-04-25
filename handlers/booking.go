package handlers

import (
	"log"
	"strconv"
	"time"

	"github.com/ConDavison1/rise-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// BookingReq is the request structure for creating a booking
// It contains the program ID and username fields that are required for booking a program.
type BookingReq struct {
	ProgramID string `json:"program_id" validate:"required"`
	Username  string `json:"user_name" validate:"required"`
}

// CreateBooking creates a new booking in the database for a specific program and user.
func CreateBooking(c *fiber.Ctx) error {
	var req BookingReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Checking to see if the program exists in the database and has enough spaces available
	var capacity, registered int
	err := db.DB.QueryRow(c.Context(), "SELECT capacity, registered FROM programs WHERE id = $1", req.ProgramID).Scan(&capacity, &registered)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Program not found",
		}) // Check if the program exists in the database
	}
	if registered >= capacity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No spaces available",
		}) // Check if the program is full
	}

	// Create a new booking
	bookingID := uuid.New().String()
	_, err = db.DB.Exec(c.Context(), "INSERT INTO bookings (id, program_id, user_name, created_at) VALUES ($1, $2, $3, $4)", bookingID, req.ProgramID, req.Username, time.Now())

	if err != nil {
		log.Println("INSERT ERROR:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create booking"})
	}

	// Update registered count in the programs table
	_, err = db.DB.Exec(c.Context(), "UPDATE programs SET registered = registered + 1 WHERE id = $1", req.ProgramID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update program",
		})
	}
	// return the successful booking
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Booking created successfully",
	})

}

type Booking struct {
	ID        string    `json:"id"`
	ProgramID string    `json:"program_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

// GetBookings retrieves all bookings from the database with pagination support.
// It returns a JSON response containing the bookings, the current page, and the number of results.
func GetBookings(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit
	// Validate page and limit values
	rows, err := db.DB.Query(
		c.Context(),
		`SELECT id, program_id, user_name, created_at FROM bookings ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	// Check for errors in the query execution
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch bookings"})
	}
	// Close the rows after processing
	defer rows.Close()
	// Check if there are any rows returned
	var bookings []Booking
	for rows.Next() {
		var b Booking
		if err := rows.Scan(&b.ID, &b.ProgramID, &b.UserName, &b.CreatedAt); err != nil {
			log.Println("SCAN ERROR (GetBookings):", err)
			continue
		}
		bookings = append(bookings, b)
	}
	log.Printf("Total bookings retrieved: %d\n", len(bookings))
	// return the bookings in JSON format
	return c.JSON(fiber.Map{
		"page":     page,
		"limit":    limit,
		"results":  len(bookings),
		"bookings": bookings,
	})
}
