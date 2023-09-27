package server

import (
	"api-tabungan/delivery/container"
	"api-tabungan/domain/health"
	"api-tabungan/domain/tabungan"
)

type handler struct {
	healthHandler   health.HealthHandler
	tabunganHandler tabungan.TabunganHandler
}

func SetupHandler(container container.Container) handler {
	return *&handler{
		healthHandler:   health.NewHealthHandler(container.HealthFeature),
		tabunganHandler: tabungan.NewTabunganHandler(container.TabunganFeature),
	}
}
