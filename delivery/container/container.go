package container

import (
	"api-tabungan/config"
	health_feature "api-tabungan/domain/health/feature"
	health_repository "api-tabungan/domain/health/repository"
	"api-tabungan/infrastructure/database"
	"api-tabungan/infrastructure/logger"
	"fmt"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	HealthFeature     health_feature.HealthFeature
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

	fmt.Println("Loading feature's...")
	// initiate feature below here!
	healthFeature := health_feature.NewHealthFeature(cfg, healthRepository)

	return Container{
		EnvironmentConfig: cfg,
		HealthFeature:     healthFeature,
	}
}
