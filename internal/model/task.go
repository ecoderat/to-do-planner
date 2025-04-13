package model

type Task struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"column:name"`
	Duration   int    `gorm:"column:duration"`
	Difficulty int    `gorm:"column:difficulty"`

	ProviderName string   `gorm:"column:provider_name"`
	Provider     Provider `gorm:"foreignKey:ProviderName;references:Name"`
}
