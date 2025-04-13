package domain

import "to-do-planner/internal/scheduler"

type Task struct {
	Name         string
	Duration     int
	Difficulty   int
	ProviderName string
}

type Tasks []Task

func (task Task) ToSchedularTask() scheduler.Task {
	return scheduler.Task{
		Name:       task.Name,
		Duration:   task.Duration,
		Difficulty: task.Difficulty,
	}
}

func (tasks Tasks) ToSchedularTasks() []scheduler.Task {
	result := make([]scheduler.Task, len(tasks))

	for i, task := range tasks {
		result[i] = task.ToSchedularTask()
	}

	return result
}
