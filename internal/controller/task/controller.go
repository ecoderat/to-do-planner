package task

import (
	"github.com/gofiber/fiber/v2"

	"to-do-planner/internal/config"
	"to-do-planner/internal/service/provider"
	"to-do-planner/internal/service/task"
)

type TaskController interface {
	GetTasks(*fiber.Ctx) error
	LoadTasks(*fiber.Ctx) error
	ScheduleTasks(*fiber.Ctx) error
}

type taskController struct {
	taskService     task.Service
	providerService provider.Service
}

func New(cfg config.Config) TaskController {
	return &taskController{
		taskService:     cfg.Services.Task,
		providerService: cfg.Services.Provider,
	}
}

func (ctrl *taskController) GetTasks(c *fiber.Ctx) error {
	tasks, err := ctrl.taskService.GetTasks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching providers")
	}

	return c.JSON(tasks)
}

func (ctrl *taskController) LoadTasks(c *fiber.Ctx) error {
	err := ctrl.taskService.LoadTasks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading tasks")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (ctrl *taskController) ScheduleTasks(c *fiber.Ctx) error {
	schedule, err := ctrl.taskService.ScheduleTasks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error scheduling tasks")
	}

	return c.JSON(schedule)
}
