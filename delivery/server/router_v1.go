package server

import (
	"api-tabungan/delivery/server/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func routerGroupV1(app *fiber.App, handler handler) {

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,OPTION",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(middleware.Logger())

	v1 := app.Group("/v1")
	{
		v1.Get("/ping", handler.healthHandler.Ping)
		v1.Get("/health-check", handler.healthHandler.HealthCheck)
		tools := v1.Group("/tools")
		{
			log := tools.Group("/log")
			{
				log.Use(middleware.Logging())
				log.Post("/xid", handler.healthHandler.GetLogDataByXID)
				log.Get("/all", handler.healthHandler.GetLogData)
			}
		}

		tabungan := v1.Group("/tabungan")
		{
			tabungan.Post("/daftar", handler.tabunganHandler.RegisterHandler)
		}
	}
}
