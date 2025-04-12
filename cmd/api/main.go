package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"to-do-planner/internal/db"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, To-Do Planner!")
	})

	log.Fatal(app.Listen(":3000")) // Listen on port 3000
}
