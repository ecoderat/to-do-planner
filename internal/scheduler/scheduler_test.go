package scheduler

import (
	"fmt"
	"testing"
)

func PrintSchedule(schedule []ScheduleSlot) {
	grouped := make(map[int][]ScheduleSlot)
	for _, slot := range schedule {
		grouped[slot.Week] = append(grouped[slot.Week], slot)
	}

	for week := 1; ; week++ {
		slots, ok := grouped[week]
		if !ok {
			break
		}
		fmt.Printf("\nWeek %d\n", week)
		fmt.Println("--------")
		for _, slot := range slots {
			fmt.Printf("%s (used %d units):\n", slot.Developer.Name, slot.LoadUsed)
			for _, task := range slot.Tasks {
				fmt.Printf("  - %s (%dh, %dx → %d units)\n", task.Name, task.Duration, task.Difficulty, task.Workload)
			}
		}
	}
}

func TestScheduleTasks(t *testing.T) {
	devs := []Developer{
		{ID: 1, Name: "DEV1", Capacity: 1},
		{ID: 2, Name: "DEV2", Capacity: 2},
		{ID: 3, Name: "DEV3", Capacity: 3},
		{ID: 4, Name: "DEV4", Capacity: 4},
		{ID: 5, Name: "DEV5", Capacity: 5},
	}

	tasks := []Task{
		{Name: "Task A", Duration: 10, Difficulty: 4},
		{Name: "Task B", Duration: 9, Difficulty: 8},
		{Name: "Task C", Duration: 17, Difficulty: 3},
		{Name: "Task D", Duration: 25, Difficulty: 2},
		{Name: "Task E", Duration: 16, Difficulty: 4},
		{Name: "Task F", Duration: 8, Difficulty: 7},
		{Name: "Task G", Duration: 10, Difficulty: 5},
		{Name: "Task H", Duration: 14, Difficulty: 4},
		{Name: "Task I", Duration: 12, Difficulty: 7},
		{Name: "Task J", Duration: 13, Difficulty: 3},
		{Name: "Task K", Duration: 14, Difficulty: 5},
		{Name: "Task L", Duration: 5, Difficulty: 3},
	}

	scheduler := NewScheduler()
	schedule := scheduler.ScheduleTasks(tasks, devs)
	PrintSchedule(schedule)
	t.FailNow()
}
