package model

type Task struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"column:name"`
	Duration   int    `gorm:"column:duration"`
	Difficulty int    `gorm:"column:difficulty"`

	ProviderID uint     `gorm:"column:provider_id"`    // FK to Provider.ID
	Provider   Provider `gorm:"foreignKey:ProviderID"` // Proper GORM relation
}
