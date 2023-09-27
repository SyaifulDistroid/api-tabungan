package response

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"api-tabungan/domain/shared/constant"
	Shared "api-tabungan/domain/shared/context"
	Error "api-tabungan/domain/shared/error"
	"api-tabungan/infrastructure/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Response struct {
	Status       string      `json:"status"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data"`
	XID          interface{} `json:"xid,omitempty"`
	ResponseCode int         `json:"code"`
}

type ErrorValidation struct {
	Status  string   `json:"status"`
	Message []string `json:"message,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

func ResponseOK(c *fiber.Ctx, msg string, data interface{}) error {

	var (
		code = http.StatusOK
	)

	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.SUCCESS,
		Message:      msg,
		XID:          xid,
		Data:         data,
		ResponseCode: code,
	}

	logger.LogInfoResponse(c.UserContext(), constant.RESPONSE, msg, response)

	return c.Status(code).JSON(response)
}

func ResponseErrorWithContext(ctx context.Context, err error) error {

	var (
		errType    string
		statusCode = http.StatusBadRequest
	)

	errType, err = Error.TrimMesssage(err)
	// Set Status Code
	if errType == constant.ErrDatabase || errType == constant.ErrTimeout {
		statusCode = http.StatusInternalServerError
	} else if errType == constant.ErrDataNotFound {
		statusCode = http.StatusNotFound
	} else if errType == constant.ErrPanic {
		statusCode = http.StatusNotImplemented
	}

	logger.LogError(ctx, constant.RESPONSE, errType, err.Error())

	errSplit := strings.Split(err.Error(), ":")
	errMessage := errSplit[0]
	if strings.Contains(err.Error(), "sql:") || errType == constant.ErrDatabase {
		sqlErr := strings.Join(errSplit[1:], "")
		errMessage = strings.TrimSpace(sqlErr)
	} else if len(errSplit) > 1 {
		errMessage = errSplit[1]
	}

	errData := ErrorMessage{
		Message: strings.TrimSpace(errMessage),
	}

	c := Shared.GetValueFiberFromContext(ctx)

	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.ERROR,
		Message:      errType,
		XID:          xid,
		Data:         errData,
		ResponseCode: statusCode,
	}

	logger.LogInfoResponse(c.UserContext(), constant.RESPONSE, err.Error(), response)

	return c.Status(statusCode).JSON(response)
}

func ResponseCustomError(c *fiber.Ctx, statusCode int, msg string, err error) error {

	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.ERROR,
		Message:      fmt.Sprintf("%s: %s", msg, err.Error()),
		Data:         nil,
		XID:          xid,
		ResponseCode: statusCode,
	}

	return c.Status(statusCode).JSON(response)
}

func ResponseValidation(c *fiber.Ctx, statusCode int, msg string, dataErr interface{}) error {
	xid := c.UserContext().Value(constant.Xid).(interface{})
	if xid == "" {
		xid = uuid.New().String()
	}

	response := Response{
		Status:       constant.ERROR,
		Message:      msg,
		XID:          xid,
		Data:         dataErr,
		ResponseCode: statusCode,
	}

	return c.Status(statusCode).JSON(response)
}
