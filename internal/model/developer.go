package model

type Developer struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"` // Tasks per hour (1x, 2x, 3x, etc.)
}
