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

	router := app.Group("/api")

	router.Use(cors)

	providerController := providerctrl.New(*cfg)
	taskController := taskctrl.New(*cfg)

	app.Get("/provider", providerController.GetProviders)
	app.Post("/provider", providerController.CreateProvider)

	app.Get("/task", taskController.GetTasks)
	app.Get("/task-load", taskController.LoadTasks)

	app.Get("/schedule", taskController.ScheduleTasks)
}
