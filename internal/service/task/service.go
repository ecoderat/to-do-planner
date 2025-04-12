package task

import (
	"context"

	taskrepo "to-do-planner/internal/repository/task"
)

type Service interface {
	GetTasks(ctx context.Context) ([]Task, error)
}

type taskService struct {
	repository taskrepo.Repository
}

func New(repository taskrepo.Repository) Service {
	return &taskService{
		repository: repository,
	}
}

type Task struct {
	Name         string
	Duration     int
	Difficulty   int
	ProviderName string
}

func (s *taskService) GetTasks(ctx context.Context) ([]Task, error) {
	tasks, err := s.repository.GetTasks(ctx)
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
