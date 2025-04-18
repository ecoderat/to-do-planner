package provider

import (
	"context"
	"encoding/json"

	"to-do-planner/internal/domain"
	"to-do-planner/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetProviders(ctx context.Context) ([]domain.ProviderConfig, error)
	GetProvider(ctx context.Context, name string) (domain.ProviderConfig, error)
	Create(ctx context.Context, provider domain.ProviderConfig) error
}

type ProviderRepository struct {
	db *gorm.DB
}

// compile time interface checks
var _ Repository = new(ProviderRepository)

func NewProviderRepository(db *gorm.DB) *ProviderRepository {
	return &ProviderRepository{db: db}
}

func (r *ProviderRepository) Create(ctx context.Context, provider domain.ProviderConfig) error {
	var modelProvider model.Provider
	modelProvider.Name = provider.ProviderName
	modelProvider.APIURL = provider.APIURL

	responseKeys, err := json.Marshal(provider.ResponseKeys)
	if err != nil {
		return err
	}
	modelProvider.ResponseKeys = string(responseKeys)

	err = r.db.WithContext(ctx).Create(&modelProvider).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProviderRepository) GetProviders(ctx context.Context) ([]domain.ProviderConfig, error) {
	var providers []model.Provider

	err := r.db.WithContext(ctx).
		Model(&model.Provider{}).
		Find(&providers).
		Error
	if err != nil {
		return nil, err
	}

	return convertToDomainProviderConfigs(providers), nil
}

func (r *ProviderRepository) GetProvider(ctx context.Context, name string) (domain.ProviderConfig, error) {
	var provider model.Provider

	err := r.db.WithContext(ctx).
		Model(&model.Provider{}).
		Where("name = ?", name).
		First(&provider).
		Error
	if err != nil {
		return domain.ProviderConfig{}, err
	}

	return convertToDomainProviderConfig(provider), nil
}

func convertToDomainResponseKeys(provider model.Provider) domain.ResponseKeys {
	var keys domain.ResponseKeys
	err := json.Unmarshal([]byte(provider.ResponseKeys), &keys)
	if err != nil {
		return domain.ResponseKeys{}
	}
	return keys
}

func convertToDomainProviderConfig(providers model.Provider) domain.ProviderConfig {
	return domain.ProviderConfig{
		ProviderName: providers.Name,
		APIURL:       providers.APIURL,
		ResponseKeys: convertToDomainResponseKeys(providers),
	}
}

func convertToDomainProviderConfigs(providers []model.Provider) []domain.ProviderConfig {
	var domainProviders []domain.ProviderConfig
	for _, provider := range providers {
		domainProviders = append(domainProviders, convertToDomainProviderConfig(provider))
	}
	return domainProviders
}
