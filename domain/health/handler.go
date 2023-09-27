package health

import (
	"api-tabungan/domain/health/feature"
	"api-tabungan/domain/health/model"

	"api-tabungan/domain/shared/constant"
	shared_constant "api-tabungan/domain/shared/constant"
	Error "api-tabungan/domain/shared/error"
	"api-tabungan/domain/shared/response"
	"api-tabungan/domain/shared/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler interface {
	Ping(c *fiber.Ctx) error
	HealthCheck(c *fiber.Ctx) error
	GetLogDataByXID(c *fiber.Ctx) error
	GetLogData(c *fiber.Ctx) error
}

type healthHandler struct {
	healthFeature feature.HealthFeature
}

func NewHealthHandler(healthFeature feature.HealthFeature) HealthHandler {
	return &healthHandler{
		healthFeature: healthFeature,
	}
}

func (hh healthHandler) Ping(c *fiber.Ctx) error {
	return response.ResponseOK(c, constant.SUCCESS, "pong!")
}

func (hh healthHandler) HealthCheck(c *fiber.Ctx) error {
	ctx := c.UserContext()

	resp, err := hh.healthFeature.HealthCheck(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}

func (hh healthHandler) GetLogDataByXID(c *fiber.Ctx) error {
	ctx := c.UserContext()

	request := new(model.LogRequest)
	if err := c.BodyParser(request); err != nil {
		if err.Error() != "Unprocessable Entity" {
			return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		}

		err = Error.New(ctx, shared_constant.ErrInvalidRequest, shared_constant.ErrInvalidRequest, errors.New(fmt.Sprintf(shared_constant.ErrInvalidBody, err.Error())))
		return response.ResponseErrorWithContext(ctx, err)
	}

	errMessage := validator.ValidateStruct(request)
	if errMessage != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, shared_constant.ErrInvalidRequest, errMessage)
	}

	resp, err := hh.healthFeature.GetLogByXID(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}

func (hh healthHandler) GetLogData(c *fiber.Ctx) error {
	ctx := c.UserContext()

	resp, err := hh.healthFeature.GetLogAll(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
