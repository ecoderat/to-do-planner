package task

import (
	"context"

	"to-do-planner/internal/domain"
	"to-do-planner/internal/provider"
	taskrepo "to-do-planner/internal/repository/task"
	providerservice "to-do-planner/internal/service/provider"
)

type Service interface {
	GetTasks(ctx context.Context) ([]Task, error)
	LoadTasks(ctx context.Context) error
}

type taskService struct {
	taskRepository  taskrepo.Repository
	providerService providerservice.Service
}

func New(taskRepository taskrepo.Repository, providerService providerservice.Service) Service {
	return &taskService{
		taskRepository:  taskRepository,
		providerService: providerService,
	}
}

type Task struct {
	Name         string
	Duration     int
	Difficulty   int
	ProviderName string
}

func (s *taskService) GetTasks(ctx context.Context) ([]Task, error) {
	tasks, err := s.taskRepository.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	var taskList []Task
	for _, t := range tasks {
		taskList = append(taskList, Task{
			Name:         t.Name,
			Duration:     t.Duration,
			Difficulty:   t.Difficulty,
			ProviderName: t.ProviderName,
		})
	}

	return taskList, nil
}

func (s *taskService) LoadTasks(ctx context.Context) error {
	providers, err := s.providerService.GetProviders(ctx)
	if err != nil {
		return err
	}

	var tasks []domain.Task
	for _, p := range providers {
		provider := provider.New(&domain.ProviderConfig{
			ProviderName: p.ProviderName,
			APIURL:       p.APIURL,
			ResponseKeys: domain.ResponseKeys{
				KeyOfTaskList:    p.ResponseKeys.KeyOfTaskList,
				HasKeyOfTaskList: p.ResponseKeys.HasKeyOfTaskList,
				TaskNameField:    p.ResponseKeys.TaskNameField,
				DurationField:    p.ResponseKeys.DurationField,
				DifficultyField:  p.ResponseKeys.DifficultyField,
			},
		})

		fetchedTasks, err := provider.FetchTasks(ctx)
		if err != nil {
			return err
		}

		tasks = append(tasks, fetchedTasks...)
	}

	// Delete all existing tasks before inserting new ones
	err = s.taskRepository.DeleteAllTasks(ctx)
	if err != nil {
		return err
	}

	// Insert the new tasks into the database
	err = s.taskRepository.CreateTasks(ctx, tasks)
	if err != nil {
		return err
	}

	return nil
}
