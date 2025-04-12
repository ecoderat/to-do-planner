package custom

import (
	"context"

	"to-do-planner/internal/domain"
)

type CustomProvider struct{}

func (f *CustomProvider) Name() string {
	return "CustomProvider"
}

func (f *CustomProvider) FetchTasks(ctx context.Context) ([]domain.Task, error) {
	// Custom logic goes here...
	return []domain.Task{
		{Name: "Custom Task", Duration: 10, Difficulty: 5},
	}, nil
}
