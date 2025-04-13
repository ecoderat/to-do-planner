package developer

import (
	"context"

	"gorm.io/gorm"

	"to-do-planner/internal/domain"
	"to-do-planner/internal/model"
)

type Repository interface {
	GetDevelopers(ctx context.Context) (domain.Developers, error)
}

type DeveloperRepository struct {
	db *gorm.DB
}

// compile time interface checks
var _ Repository = new(DeveloperRepository)

func New(db *gorm.DB) *DeveloperRepository {
	return &DeveloperRepository{db: db}
}

func (r *DeveloperRepository) GetDevelopers(ctx context.Context) (domain.Developers, error) {
	var developers domain.Developers

	err := r.db.WithContext(ctx).
		Model(&model.Developer{}).
		Select("name, capacity").
		Find(&developers).
		Error
	if err != nil {
		return nil, err
	}

	return developers, nil
}
