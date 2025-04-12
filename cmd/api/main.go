package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"to-do-planner/internal/db"
	"to-do-planner/internal/domain"
	providerint "to-do-planner/internal/provider/mapped"
	providerrepo "to-do-planner/internal/repository/provider"
	taskrepo "to-do-planner/internal/repository/task"
)

func main() {
	// Initialize the database connection
	db.InitDB()

	app := fiber.New()

	providerRepo := providerrepo.NewProviderRepository(db.DB)
	taskRepo := taskrepo.NewTaskRepository(db.DB)

	app.Get("/provider", func(c *fiber.Ctx) error {
		log.Println("Fetching providers...")
		// fetch the providers from the database
		providers, err := providerRepo.GetProviders(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching providers")
		}

		// return the providers as JSON
		return c.JSON(providers)
	})

	app.Get("/task", func(c *fiber.Ctx) error {
		log.Println("Fetching tasks...")
		// fetch the tasks from the database
		tasks, err := taskRepo.GetTasks(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching providers")
		}

		// return the tasks as JSON
		return c.JSON(tasks)
	})

	app.Get("/task-load", func(c *fiber.Ctx) error {
		log.Println("Loading tasks from apis...")

		// get all providers from the database
		providers, err := providerRepo.GetProviders(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching providers")
		}

		var tasks []domain.Task
		for _, provider := range providers {
			fmt.Printf("provider: %+v\n", provider)
			// create a new mapped provider
			mappedProvider := providerint.NewMappedProvider(&provider)
			// fetch tasks from the provider
			fetchedTasks, err := mappedProvider.FetchTasks(c.Context())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error fetching tasks from provider")
			}
			// append the fetched tasks to the tasks slice
			tasks = append(tasks, fetchedTasks...)
		}

		// return the tasks as JSON
		return c.JSON(tasks)
	})

	app.Post("/provider", func(c *fiber.Ctx) error {
		log.Println("Creating provider...")
		// create a new provider
		var provider domain.ProviderConfig
		if err := c.BodyParser(&provider); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		// save the provider to the database
		err := providerRepo.Create(c.Context(), provider)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error creating provider")
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	log.Fatal(app.Listen(":3000")) // Listen on port 3000
}
