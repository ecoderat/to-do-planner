package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"to-do-planner/internal/config"
	providerctrl "to-do-planner/internal/controller/provider"
	taskctrl "to-do-planner/internal/controller/task"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config) {
	cors := cors.New()

	router := app.Group("/")

	router.Use(cors)

	providerController := providerctrl.New(*cfg)
	taskController := taskctrl.New(*cfg)

	router.Get("/provider", providerController.GetProviders)
	router.Post("/provider", providerController.CreateProvider)

	router.Get("/task", taskController.GetTasks)
	router.Get("/task-load", taskController.LoadTasks)

	router.Get("/schedule", taskController.ScheduleTasks)
}
