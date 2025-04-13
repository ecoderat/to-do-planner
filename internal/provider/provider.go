package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"github.com/tidwall/gjson"

	"to-do-planner/internal/domain"
)

type Provider struct {
	Config *domain.ProviderConfig
	Client *http.Client
}

func New(config *domain.ProviderConfig) domain.Provider {
	return &Provider{
		Config: config,
		Client: &http.Client{},
	}
}

func (mp *Provider) Name() string {
	return mp.Config.ProviderName
}

func (mp *Provider) FetchTasks(ctx context.Context) ([]domain.Task, error) {
	log.Infof("Fetching tasks from %s", mp.Config.APIURL)
	// Create a new HTTP request with the context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, mp.Config.APIURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := mp.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tasks []domain.Task
	result := gjson.ParseBytes(body)

	taskItems := result
	if mp.Config.ResponseKeys.HasKeyOfTaskList {
		taskItems = result.Get(mp.Config.ResponseKeys.KeyOfTaskList)
	}

	if !taskItems.IsArray() {
		return nil, fmt.Errorf("invalid task list")
	}

	for _, item := range taskItems.Array() {
		task := domain.Task{
			Name:         item.Get(mp.Config.ResponseKeys.TaskNameField).String(),
			Duration:     int(item.Get(mp.Config.ResponseKeys.DurationField).Int()),
			Difficulty:   int(item.Get(mp.Config.ResponseKeys.DifficultyField).Int()),
			ProviderName: mp.Config.ProviderName,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
