package task

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"to-do-planner/internal/config"
	"to-do-planner/internal/domain"
	providerint "to-do-planner/internal/provider/mapped"
	"to-do-planner/internal/service/provider"
	"to-do-planner/internal/service/task"
)

type TaskController interface {
	GetTasks(*fiber.Ctx) error
	LoadTasks(*fiber.Ctx) error
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
	log.Println("Loading tasks from apis...")

	providers, err := ctrl.providerService.GetProviders(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching providers")
	}

	var tasks []domain.Task
	for _, provider := range providers {
		mappedProvider := providerint.NewMappedProvider(&domain.ProviderConfig{
			ProviderName: provider.ProviderName,
			APIURL:       provider.APIURL,
			ResponseKeys: domain.ResponseKeys{
				KeyOfTaskList:    provider.ResponseKeys.KeyOfTaskList,
				HasKeyOfTaskList: provider.ResponseKeys.HasKeyOfTaskList,
				TaskNameField:    provider.ResponseKeys.TaskNameField,
				DurationField:    provider.ResponseKeys.DurationField,
				DifficultyField:  provider.ResponseKeys.DifficultyField,
			},
		})

		fetchedTasks, err := mappedProvider.FetchTasks(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching tasks from provider")
		}

		tasks = append(tasks, fetchedTasks...)
	}

	return c.JSON(tasks)
}
