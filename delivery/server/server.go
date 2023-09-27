package server

import (
	"fmt"
	"net/http"

	"api-tabungan/delivery/container"
	shared_constant "api-tabungan/domain/shared/constant"
	"api-tabungan/domain/shared/context"
	Error "api-tabungan/domain/shared/error"
	"api-tabungan/domain/shared/response"

	"github.com/gofiber/fiber/v2"
)

func ServeHttp(cont container.Container) *fiber.App {
	fmt.Println("Starting http service...")

	handler := SetupHandler(cont)

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				ctx := context.CreateContext()
				ctx = context.SetFiberToContext(ctx, c)

				err = Error.New(ctx, shared_constant.ErrPanic, shared_constant.ErrPanicWhenExecuteAPI, err)

				return response.ResponseErrorWithContext(ctx, err)
			},
		},
	)

	// iniate router v1
	routerGroupV1(app, handler)

	app.Use(func(c *fiber.Ctx) error {
		resp := fiber.Map{
			"status":  fmt.Sprintf("route %s or method not allowed", http.StatusText(http.StatusNotFound)),
			"message": fmt.Sprintf("route %s", http.StatusText(http.StatusNotFound)),
		}
		return c.Status(fiber.StatusNotFound).JSON(resp)
	})

	return app
}
