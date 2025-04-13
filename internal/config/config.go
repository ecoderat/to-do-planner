package config

import (
	"gorm.io/gorm"

	providerrepo "to-do-planner/internal/repository/provider"
	taskrepo "to-do-planner/internal/repository/task"
	providerservice "to-do-planner/internal/service/provider"
	taskservice "to-do-planner/internal/service/task"
)

type Config struct {
	Services     Services
	Repositories Repositories
}

type Services struct {
	Provider providerservice.Service
	Task     taskservice.Service
}

type Repositories struct {
	Provider providerrepo.Repository
	Task     taskrepo.Repository
}

func Load(db *gorm.DB) *Config {
	cfg := &Config{}
	providerRepo := providerrepo.NewProviderRepository(db)
	taskRepo := taskrepo.NewTaskRepository(db)

	cfg.Repositories = Repositories{
		Provider: providerRepo,
		Task:     taskRepo,
	}

	providerService := providerservice.New(cfg.Repositories.Provider)
	taskService := taskservice.New(cfg.Repositories.Task, providerService)

	cfg.Services = Services{
		Provider: providerService,
		Task:     taskService,
	}

	return cfg
}
