package main

import (
	"flash/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Serve static files (like CSS) from the /static directory
	app.Static("/css", "./static/css")

	// Serve the main index.html file
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./views/index.html")
	})

	// Define the API endpoint for performing the speed test
	app.Get("/api/speedtest", handlers.SpeedTestHandler)

	// Start the server on port 8080
	app.Listen(":8080")
}
