package model

type Developer struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name;unique"`
	Capacity int    `gorm:"column:capacity"` // Tasks per hour (1x, 2x, 3x, etc.)
}
