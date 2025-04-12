package domain

import "context"

type ResponseKeys struct {
	HasKeyOfTaskList bool   `json:"hasKeyOfTaskList"`
	KeyOfTaskList    string `json:"keyOfTaskList"`
	TaskNameField    string `json:"taskNameField"`
	DurationField    string `json:"durationField"`
	DifficultyField  string `json:"difficultyField"`
}

type ProviderConfig struct {
	ProviderName string       `json:"providerName"`
	APIURL       string       `json:"apiURL"`
	ResponseKeys ResponseKeys `json:"responseKeys"`
}

type Provider interface {
	FetchTasks(ctx context.Context) ([]Task, error)
	Name() string
}
