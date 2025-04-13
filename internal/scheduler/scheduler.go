package scheduler

import (
	"sort"
)

const MaxWeeklyHours = 45 // 5 days * 9 hours

type Scheduler interface {
	ScheduleTasks(tasks []Task, devs []Developer) []ScheduleSlot
}

type scheduler struct{}

func NewScheduler() Scheduler {
	return &scheduler{}
}

type Task struct {
	Name       string
	Duration   int // in hours
	Difficulty int // 1x–5x
	Workload   int // Computed: Duration * Difficulty
}

type Developer struct {
	ID       uint
	Name     string
	Capacity int // How many "x" they can complete per hour
}

type ScheduleSlot struct {
	Week      int
	Developer Developer
	Tasks     []Task
	LoadUsed  int
}

func (s *scheduler) ScheduleTasks(tasks []Task, developers []Developer) []ScheduleSlot {
	// Calculate workload for each task
	for i := range tasks {
		tasks[i].Workload = tasks[i].Duration * tasks[i].Difficulty
	}

	// Sort tasks by descending workload
	sort.Slice(developers, func(i, j int) bool {
		return developers[i].Capacity > developers[j].Capacity
	})

	var schedule []ScheduleSlot
	week := 1

	for len(tasks) > 0 {
		for _, dev := range developers {
			devWeeklyCapacity := dev.Capacity * MaxWeeklyHours
			used := 0
			assigned := []Task{}

			// Try to assign tasks to this developer
			for i := 0; i < len(tasks); {
				if tasks[i].Workload <= (devWeeklyCapacity - used) {
					assigned = append(assigned, tasks[i])
					used += tasks[i].Workload
					// Remove task
					tasks = append(tasks[:i], tasks[i+1:]...)
				} else {
					i++
				}
			}

			if len(assigned) > 0 {
				schedule = append(schedule, ScheduleSlot{
					Week:      week,
					Developer: dev,
					Tasks:     assigned,
					LoadUsed:  used,
				})
			}
		}
		week++
	}

	return schedule
}
