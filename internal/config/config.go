package config

import (
	"gorm.io/gorm"

	developerrepo "to-do-planner/internal/repository/developer"
	providerrepo "to-do-planner/internal/repository/provider"
	taskrepo "to-do-planner/internal/repository/task"
	"to-do-planner/internal/scheduler"
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
	Provider  providerrepo.Repository
	Task      taskrepo.Repository
	Developer developerrepo.Repository
}

func Load(db *gorm.DB) *Config {
	cfg := &Config{}
	providerRepo := providerrepo.NewProviderRepository(db)
	taskRepo := taskrepo.NewTaskRepository(db)
	developerRepo := developerrepo.New(db)

	scheduler := scheduler.NewScheduler()

	cfg.Repositories = Repositories{
		Provider:  providerRepo,
		Task:      taskRepo,
		Developer: developerRepo,
	}

	providerService := providerservice.New(cfg.Repositories.Provider)
	taskService := taskservice.New(cfg.Repositories.Task, providerService, cfg.Repositories.Developer, scheduler)

	cfg.Services = Services{
		Provider: providerService,
		Task:     taskService,
	}

	return cfg
}
