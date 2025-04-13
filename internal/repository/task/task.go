package task

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"to-do-planner/internal/domain"
	"to-do-planner/internal/model"
)

type Repository interface {
	GetTasks(ctx context.Context) (domain.Tasks, error)
	CreateTasks(ctx context.Context, tasks domain.Tasks) error
	DeleteAllTasks(ctx context.Context) error
}

type TaskRepository struct {
	db *gorm.DB
}

// compile time interface checks
var _ Repository = new(TaskRepository)

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks(ctx context.Context) (domain.Tasks, error) {
	var result domain.Tasks

	err := r.db.WithContext(ctx).
		Table("tasks").
		Select("tasks.id, tasks.name, tasks.duration, tasks.difficulty, providers.name as provider_name").
		Joins("left join providers on providers.name = tasks.provider_name").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *TaskRepository) CreateTasks(ctx context.Context, tasks domain.Tasks) error {
	if len(tasks) == 0 {
		return nil
	}

	modelTasks := make([]model.Task, len(tasks))
	for i, task := range tasks {
		modelTasks[i] = model.Task{
			Name:         task.Name,
			Duration:     task.Duration,
			Difficulty:   task.Difficulty,
			ProviderName: task.ProviderName,
		}
	}

	err := r.db.WithContext(ctx).
		Create(&modelTasks).Error
	if err != nil {
		fmt.Println("Create error:", err)
		return err
	}

	return nil
}

func (r *TaskRepository) DeleteAllTasks(ctx context.Context) error {
	err := r.db.Debug().Exec(`DELETE FROM "tasks"`).Error
	if err != nil {
		fmt.Println("Delete error:", err)
		return err
	}
	return nil
}
