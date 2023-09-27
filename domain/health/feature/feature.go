package feature

import (
	"api-tabungan/config"
	"api-tabungan/domain/health/constant"
	"api-tabungan/domain/health/helper"
	"api-tabungan/domain/health/model"
	"api-tabungan/domain/health/repository"
	shared_constant "api-tabungan/domain/shared/constant"
	Error "api-tabungan/domain/shared/error"
	"api-tabungan/infrastructure/logger"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"
)

type HealthFeature interface {
	HealthCheck(ctx context.Context) (resp model.HealthCheck, err error)
	GetLogByXID(ctx context.Context, request *model.LogRequest) (resp *model.LogDetailResponse, err error)
	GetLogAll(ctx context.Context) (resp *model.LogDetailResponse, err error)
}

type healthFeature struct {
	config           config.EnvironmentConfig
	healthRepository repository.HealthRepository
}

func NewHealthFeature(config config.EnvironmentConfig, healthRepository repository.HealthRepository) HealthFeature {
	return &healthFeature{
		config:           config,
		healthRepository: healthRepository,
	}
}

func (hf healthFeature) HealthCheck(ctx context.Context) (resp model.HealthCheck, err error) {

	var (
		status   = constant.HEALTHY
		dbstatus = constant.CONNECTED
	)

	db, err := hf.healthRepository.DatabaseHealth(ctx)
	if !db {
		status = constant.NOT_HEALTHY
		dbstatus = constant.DISCONECTED
	} else if err != nil {
		status = constant.NOT_HEALTHY
	}

	resp = model.HealthCheck{
		AppDetail: model.AppDetail{
			Name:    hf.config.App.Name,
			Version: hf.config.App.Version,
		},
		DatabaseDetail: model.DatabaseDetail{
			Dialect: hf.config.Database.Dialect,
			Status:  dbstatus,
		},
		Status: status,
	}

	return

}

func (hf healthFeature) GetLogByXID(ctx context.Context, request *model.LogRequest) (resp *model.LogDetailResponse, err error) {

	var (
		logDetail model.LogDetailResponse
	)

	uuid, err := uuid.Parse(request.Xid)
	if err != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrInvalidXID, errors.New(fmt.Sprintf(constant.ErrInvalidXIDWithError, err.Error())))
		return
	}

	logDetail.Xid = &uuid

	logs, err := helper.ReadLines(logger.LogName)
	if err != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrLogDataNotFound, errors.New(fmt.Sprintf(constant.ErrLogDataNotFoundWithError, logger.LogName, err.Error())))
		return
	}

	for _, log := range logs {
		if strings.Contains(log, uuid.String()) {
			logDetail.Contents = append(logDetail.Contents, log)
		}
	}

	resp = &logDetail

	return
}

func (hf healthFeature) GetLogAll(ctx context.Context) (resp *model.LogDetailResponse, err error) {

	var (
		logDetail model.LogDetailResponse
	)

	content, err := ioutil.ReadFile(logger.LogName)
	if err != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrLogDataNotFound, errors.New(fmt.Sprintf(constant.ErrLogDataNotFoundWithError, logger.LogName, err.Error())))
		return
	}

	logDetail.Contents = append(logDetail.Contents, string(content))

	resp = &logDetail

	return
}
