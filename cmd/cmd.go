package cmd

import (
	"api-tabungan/delivery/container"
	"api-tabungan/delivery/server"
	"fmt"
)

func Execute() {
	// define container
	container := container.SetupContainer()

	// start http service
	server := server.ServeHttp(container)
	server.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
}
