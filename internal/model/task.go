package model

type Task struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	Duration   int    `json:"duration"` // In hours
	Difficulty int    `json:"difficulty"`
	Provider   string `json:"provider"` // e.g., "Provider1", "Provider2"
}
