package server

import (
	"api-tabungan/delivery/container"
	"api-tabungan/domain/health"
)

type handler struct {
	healthHandler health.HealthHandler
}

func SetupHandler(container container.Container) handler {
	return *&handler{
		healthHandler: health.NewHealthHandler(container.HealthFeature),
	}
}
