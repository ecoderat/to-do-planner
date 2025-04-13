package domain

type ScheduleSlot struct {
	Week      int
	Developer Developer
	Tasks     []Task
	LoadUsed  int
}
