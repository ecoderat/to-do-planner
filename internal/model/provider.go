package model

type Provider struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	APIURL string `gorm:"column:api_url"`
	// ResponseKeys contains the keys used to parse the response from the provider
	ResponseKeys string `gorm:"column:response_keys"` // JSON string
	// Example: `{"HasKeyOfTaskList": true, "KeyOfTaskList": "tasks", "TaskNameField": "name", "DurationField": "duration", "DifficultyField": "difficulty"}`
}
