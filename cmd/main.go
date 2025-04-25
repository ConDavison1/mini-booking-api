package main
// Import the Fiber web framework
import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API")
	})

	app.Listen(":3000")
}
