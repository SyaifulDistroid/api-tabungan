package feature

import (
	"api-tabungan/config"
	shared_constant "api-tabungan/domain/shared/constant"
	Error "api-tabungan/domain/shared/error"
	"api-tabungan/domain/tabungan/constant"
	"api-tabungan/domain/tabungan/helper"
	"api-tabungan/domain/tabungan/model"
	"api-tabungan/domain/tabungan/repository"
	"context"
	"errors"
	"fmt"
)

type TabunganFeature interface {
	CreateRekeningFeature(ctx context.Context, request model.CreateRekeningRequest) (response model.NasabahJSON, err error)
}

type tabunganFeature struct {
	config             config.EnvironmentConfig
	tabunganRepository repository.TabunganRepository
}

func NewTabunganFeature(config config.EnvironmentConfig, tabunganRepository repository.TabunganRepository) TabunganFeature {
	return &tabunganFeature{
		config:             config,
		tabunganRepository: tabunganRepository,
	}
}

func (t *tabunganFeature) CreateRekeningFeature(ctx context.Context, request model.CreateRekeningRequest) (response model.NasabahJSON, err error) {
	// check NIK
	checkNIK, _ := t.tabunganRepository.GetByNIK(ctx, request.NIK)

	if checkNIK != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNIKAlreadyExist, errors.New(fmt.Sprintf(constant.ErrNIKAlreadyExistWithNIK, request.NIK)))
		return
	}

	// check NoHP
	checkNoHP, _ := t.tabunganRepository.GetByNoHandphone(ctx, request.NoHandphone)

	if checkNoHP != nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoHandphoneAlreadyExist, errors.New(fmt.Sprintf(constant.ErrNoHandphoneAlreadyExistWithNoHP, request.NoHandphone)))
		return
	}

	// Fill Data
	data := model.Rekening{
		Nama:        request.Nama,
		NIK:         request.NIK,
		NoHandphone: request.NoHandphone,
		NoRekening:  helper.GenerateNoRekening(request),
		Saldo:       int64(constant.DEFAULT_BALANCE),
		CreatedBy:   constant.SYSTEM,
	}

	err = t.tabunganRepository.CreateRekeningRepository(ctx, &data)
	if err != nil {
		return
	}

	response = model.NasabahJSON{
		NoRekening: data.NoRekening,
	}
	return
}
