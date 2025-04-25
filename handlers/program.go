package handlers

import (
	"log"
	"time"

	"github.com/ConDavison1/rise-api/db"
	"github.com/gofiber/fiber/v2"
)
// Program is a struct that represents a program in the database.
// It contains fields for the program ID, name, capacity, registered count, visibility status, start date, and end date.
type Program struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Capacity   int       `json:"capacity"`
	Registered int       `json:"registered"`
	Visibility string    `json:"visibility"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

// GetPrograms retrieves all programs from the database and returns them as a JSON response.
func GetPrograms(c *fiber.Ctx) error {
	var programs []Program
	rows, err := db.DB.Query(c.Context(), "SELECT id, name, capacity, registered, visibility, start_date, end_date FROM programs")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch programs"})
	}
	defer rows.Close()// Ensure rows are closed after use

	// Iterate through the rows and scan the data into the Program struct
	for rows.Next() {
		var program Program
		if err := rows.Scan(
			&program.ID, &program.Name, &program.Capacity, &program.Registered, &program.Visibility, &program.StartDate, &program.EndDate); err != nil {
			log.Println("SCAN ERROR (GetPrograms):", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not scan program"})
		}

		programs = append(programs, program)
	}

	return c.JSON(programs)// Return the list of programs as JSON
}

// GetProgramByID retrieves a program by its ID and returns the program details along with the number of spaces available.
// If the program is not found, it returns a 404 error.
func GetProgramByID(c *fiber.Ctx) error {
	programID := c.Params("id")
	var program Program
	err := db.DB.QueryRow(c.Context(), "SELECT id, name, capacity, registered, visibility, start_date, end_date FROM programs WHERE id = $1", programID).Scan(
		&program.ID, &program.Name, &program.Capacity, &program.Registered, &program.Visibility, &program.StartDate, &program.EndDate)
	if err != nil {
		log.Println("SCAN ERROR (GetProgramByID):", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Program not found"})
	}
	
	return c.JSON(fiber.Map{
		"program":          program,
		"spaces_available": program.Capacity - program.Registered,
	}) // Return the program details along with the number of spaces available

}
