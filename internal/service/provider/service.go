package provider

import (
	"context"

	"to-do-planner/internal/domain"
	providerrepo "to-do-planner/internal/repository/provider"
)

type Service interface {
	GetProviders(ctx context.Context) ([]Provider, error)
	GetProvider(ctx context.Context, name string) (Provider, error)
	Create(ctx context.Context, provider Provider) error
}

type providerService struct {
	repository providerrepo.Repository
}

func New(repository providerrepo.Repository) Service {
	return &providerService{
		repository: repository,
	}
}

type ResponseKeys struct {
	HasKeyOfTaskList bool   `json:"hasKeyOfTaskList"`
	KeyOfTaskList    string `json:"keyOfTaskList"`
	TaskNameField    string `json:"taskNameField"`
	DurationField    string `json:"durationField"`
	DifficultyField  string `json:"difficultyField"`
}

type Provider struct {
	ProviderName string       `json:"providerName"`
	APIURL       string       `json:"apiURL"`
	ResponseKeys ResponseKeys `json:"responseKeys"`
}

func (s *providerService) GetProviders(ctx context.Context) ([]Provider, error) {
	providers, err := s.repository.GetProviders(ctx)
	if err != nil {
		return nil, err
	}

	var providerList []Provider
	for _, p := range providers {
		providerList = append(providerList, Provider{
			ProviderName: p.ProviderName,
			APIURL:       p.APIURL,
			ResponseKeys: ResponseKeys{
				HasKeyOfTaskList: p.ResponseKeys.HasKeyOfTaskList,
				KeyOfTaskList:    p.ResponseKeys.KeyOfTaskList,
				TaskNameField:    p.ResponseKeys.TaskNameField,
				DurationField:    p.ResponseKeys.DurationField,
				DifficultyField:  p.ResponseKeys.DifficultyField,
			},
		})
	}

	return providerList, nil
}

func (s *providerService) GetProvider(ctx context.Context, name string) (Provider, error) {
	provider, err := s.repository.GetProvider(ctx, name)
	if err != nil {
		return Provider{}, err
	}

	if provider.ProviderName == "" {
		return Provider{}, nil
	}

	return Provider{
		ProviderName: provider.ProviderName,
		APIURL:       provider.APIURL,
		ResponseKeys: ResponseKeys{
			HasKeyOfTaskList: provider.ResponseKeys.HasKeyOfTaskList,
			KeyOfTaskList:    provider.ResponseKeys.KeyOfTaskList,
			TaskNameField:    provider.ResponseKeys.TaskNameField,
			DurationField:    provider.ResponseKeys.DurationField,
			DifficultyField:  provider.ResponseKeys.DifficultyField,
		},
	}, nil
}

func (s *providerService) Create(ctx context.Context, provider Provider) error {
	providerConfig := domain.ProviderConfig{
		ProviderName: provider.ProviderName,
		APIURL:       provider.APIURL,
		ResponseKeys: domain.ResponseKeys{
			HasKeyOfTaskList: provider.ResponseKeys.HasKeyOfTaskList,
			KeyOfTaskList:    provider.ResponseKeys.KeyOfTaskList,
			TaskNameField:    provider.ResponseKeys.TaskNameField,
			DurationField:    provider.ResponseKeys.DurationField,
			DifficultyField:  provider.ResponseKeys.DifficultyField,
		},
	}

	err := s.repository.Create(ctx, providerConfig)
	if err != nil {
		return err
	}

	return nil
}
