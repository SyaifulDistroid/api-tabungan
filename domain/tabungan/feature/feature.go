package feature

import (
	"api-tabungan/config"
	"api-tabungan/domain/tabungan/repository"
)

type TabunganFeature interface {
}

type tabunganFeature struct {
	config             config.EnvironmentConfig
	tabunganRepository repository.TabunganRepository
}

func NewTabunganFeature(config config.EnvironmentConfig, tabunganRepository repository.TabunganRepository) TabunganFeature {
	return &tabunganFeature{
		config:             config,
		tabunganRepository: tabunganRepository,
	}
}
