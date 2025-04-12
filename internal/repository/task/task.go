package task

import (
	"context"
	"to-do-planner/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	GetTasks(ctx context.Context) ([]domain.Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

// compile time interface checks
var _ Repository = new(TaskRepository)

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks(ctx context.Context) ([]domain.Task, error) {
	var result []domain.Task

	err := r.db.WithContext(ctx).
		Table("tasks").
		Select("tasks.id, tasks.name, tasks.duration, tasks.difficulty, providers.name as provider_name").
		Joins("left join providers on providers.id = tasks.provider_id").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return result, nil
}
