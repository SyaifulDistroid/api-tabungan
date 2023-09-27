package container

import (
	"api-tabungan/config"
	health_feature "api-tabungan/domain/health/feature"
	health_repository "api-tabungan/domain/health/repository"
	tabungan_feature "api-tabungan/domain/tabungan/feature"
	tabungan_repository "api-tabungan/domain/tabungan/repository"
	"api-tabungan/infrastructure/database"
	"api-tabungan/infrastructure/logger"
	"fmt"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	HealthFeature     health_feature.HealthFeature
	TabunganFeature   tabungan_feature.TabunganFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	cfg, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading logger...")
	logger.InitializeLogrusLogger(cfg.Log)

	fmt.Println("Loading database...")
	db, err := database.LoadPsqlDatabase(cfg.Database)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading repository's...")
	// initiate repository below here!
	healthRepository := health_repository.NewHealthRepository(db)
	tabunganRepository := tabungan_repository.NewTabunganRepository(db)

	fmt.Println("Loading feature's...")
	// initiate feature below here!
	healthFeature := health_feature.NewHealthFeature(cfg, healthRepository)
	tabunganFeature := tabungan_feature.NewTabunganFeature(cfg, tabunganRepository)

	return Container{
		EnvironmentConfig: cfg,
		HealthFeature:     healthFeature,
		TabunganFeature:   tabunganFeature,
	}
}
