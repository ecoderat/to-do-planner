package task

import (
	"context"

	"to-do-planner/internal/domain"
	"to-do-planner/internal/provider"
	"to-do-planner/internal/repository/developer"
	taskrepo "to-do-planner/internal/repository/task"
	"to-do-planner/internal/scheduler"
	providerservice "to-do-planner/internal/service/provider"
)

type Service interface {
	GetTasks(ctx context.Context) ([]Task, error)
	LoadTasks(ctx context.Context) error
	ScheduleTasks(ctx context.Context) ([]domain.ScheduleSlot, error)
}

type taskService struct {
	taskRepository      taskrepo.Repository
	providerService     providerservice.Service
	developerRepository developer.Repository
	scheduler           scheduler.Scheduler
}

func New(taskRepository taskrepo.Repository, providerService providerservice.Service, developerRepository developer.Repository, scheduler scheduler.Scheduler) Service {
	return &taskService{
		taskRepository:      taskRepository,
		providerService:     providerService,
		developerRepository: developerRepository,
		scheduler:           scheduler,
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

	var tasks domain.Tasks
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

func (s *taskService) ScheduleTasks(ctx context.Context) ([]domain.ScheduleSlot, error) {
	tasks, err := s.taskRepository.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	developers, err := s.developerRepository.GetDevelopers(ctx)
	if err != nil {
		return nil, err
	}

	schedules := s.scheduler.ScheduleTasks(tasks.ToSchedularTasks(), developers.ToSchedularDevelopers())

	var scheduleSlots []domain.ScheduleSlot
	for _, schedule := range schedules {
		scheduleSlots = append(scheduleSlots, domain.ScheduleSlot{
			Week: schedule.Week,
			Developer: domain.Developer{
				Name:     schedule.Developer.Name,
				Capacity: schedule.Developer.Capacity,
			},
			Tasks:    convertSchedulerTasksToDomainTasks(schedule.Tasks),
			LoadUsed: schedule.LoadUsed,
		})
	}

	return scheduleSlots, nil
}

func convertSchedulerTasksToDomainTasks(schedulerTasks []scheduler.Task) domain.Tasks {
	var domainTasks domain.Tasks
	for _, t := range schedulerTasks {
		domainTasks = append(domainTasks, domain.Task{
			Name:       t.Name,
			Duration:   t.Duration,
			Difficulty: t.Difficulty,
		})
	}
	return domainTasks
}
