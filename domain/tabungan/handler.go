package tabungan

import (
	"api-tabungan/domain/shared/constant"
	"api-tabungan/domain/shared/response"
	"api-tabungan/domain/shared/validator"
	"api-tabungan/domain/tabungan/feature"
	"api-tabungan/domain/tabungan/model"
	"fmt"
	"net/http"

	Error "api-tabungan/domain/shared/error"
	"github.com/gofiber/fiber/v2"
)

type TabunganHandler interface {
	RegisterHandler(c *fiber.Ctx) error
	SavingHandler(c *fiber.Ctx) error
	WitdrawalHandler(c *fiber.Ctx) error
	BalanceHandler(c *fiber.Ctx) error
	HistoryHandler(c *fiber.Ctx) error
}

type tabunganHandler struct {
	tabunganFeature feature.TabunganFeature
}

func NewTabunganHandler(tabunganFeature feature.TabunganFeature) TabunganHandler {
	return &tabunganHandler{
		tabunganFeature: tabunganFeature,
	}
}

func (handler tabunganHandler) RegisterHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.CreateRekeningRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.tabunganFeature.CreateRekeningFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler tabunganHandler) SavingHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.SavingRekeningRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.tabunganFeature.SavingFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler tabunganHandler) WitdrawalHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var bodyReq model.WitdrawalRekeningRequest
	err := c.BodyParser(&bodyReq)
	if err != nil {
		return response.ResponseCustomError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	errResp := validator.ValidateStruct(bodyReq)
	if errResp != nil {
		return response.ResponseValidation(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errResp)
	}

	data, err := handler.tabunganFeature.WitdrawalFeature(ctx, bodyReq)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, data)
}

func (handler tabunganHandler) BalanceHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	norek := c.Params("no_rekening")
	if norek == "" || norek == "0" {
		err := Error.New(ctx, constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrInvalidRequest))
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := handler.tabunganFeature.BalanceFeature(ctx, norek)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}

func (handler tabunganHandler) HistoryHandler(c *fiber.Ctx) error {
	ctx := c.UserContext()

	norek := c.Params("no_rekening")
	if norek == "" || norek == "0" {
		err := Error.New(ctx, constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrInvalidRequest))
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := handler.tabunganFeature.HistoryFeature(ctx, norek)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
