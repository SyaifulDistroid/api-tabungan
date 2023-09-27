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
	SavingFeature(ctx context.Context, request model.SavingRekeningRequest) (response model.SaldoJSON, err error)
	WitdrawalFeature(ctx context.Context, request model.WitdrawalRekeningRequest) (response model.SaldoJSON, err error)
	BalanceFeature(ctx context.Context, noRekening string) (response model.SaldoJSON, err error)
	HistoryFeature(ctx context.Context, noRekening string) (response []model.MutasiJSON, err error)
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

func (t *tabunganFeature) SavingFeature(ctx context.Context, request model.SavingRekeningRequest) (response model.SaldoJSON, err error) {
	// Get Data
	getNorek, _ := t.tabunganRepository.GetByNoRekening(ctx, request.NoRekening)

	if getNorek == nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, request.NoRekening)))
		return
	}

	// Fill Data
	data := model.Rekening{
		ID:         getNorek.ID,
		NoRekening: getNorek.NoRekening,
		Saldo:      getNorek.Saldo + request.Nominal,
	}

	err = t.tabunganRepository.UpdateSaldoRepository(ctx, &data)
	if err != nil {
		return
	}

	// Fill Data Mutasi
	mutasi := model.Mutasi{
		NoRekening:    data.NoRekening,
		KodeTransaksi: constant.CREDIT,
		Nominal:       request.Nominal,
		Saldo:         data.Saldo,
		CreatedBy:     constant.SYSTEM,
	}

	err = t.tabunganRepository.CreateHistoryRepository(ctx, &mutasi)
	if err != nil {
		return
	}

	response = model.SaldoJSON{
		NoRekening: data.NoRekening,
		Saldo:      data.Saldo,
	}
	return
}

func (t *tabunganFeature) WitdrawalFeature(ctx context.Context, request model.WitdrawalRekeningRequest) (response model.SaldoJSON, err error) {
	// Get Data
	getNorek, _ := t.tabunganRepository.GetByNoRekening(ctx, request.NoRekening)

	if getNorek == nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, request.NoRekening)))
		return
	}

	// Validate Saldo
	if request.Nominal > getNorek.Saldo {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrSaldo, errors.New(fmt.Sprint(constant.ErrSaldo)))
		return
	}

	// Fill Data
	data := model.Rekening{
		ID:         getNorek.ID,
		NoRekening: getNorek.NoRekening,
		Saldo:      getNorek.Saldo - request.Nominal,
	}

	err = t.tabunganRepository.UpdateSaldoRepository(ctx, &data)
	if err != nil {
		return
	}

	// Fill Data Mutasi
	mutasi := model.Mutasi{
		NoRekening:    data.NoRekening,
		KodeTransaksi: constant.DEBIT,
		Nominal:       request.Nominal,
		Saldo:         data.Saldo,
		CreatedBy:     constant.SYSTEM,
	}

	err = t.tabunganRepository.CreateHistoryRepository(ctx, &mutasi)
	if err != nil {
		return
	}

	response = model.SaldoJSON{
		NoRekening: data.NoRekening,
		Saldo:      data.Saldo,
	}
	return
}

func (t *tabunganFeature) BalanceFeature(ctx context.Context, noRekening string) (response model.SaldoJSON, err error) {
	getNorek, _ := t.tabunganRepository.GetByNoRekening(ctx, helper.StrToInt64(noRekening))

	if getNorek == nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, noRekening)))
		return
	}

	response = model.SaldoJSON{
		NoRekening: getNorek.NoRekening,
		Saldo:      getNorek.Saldo,
	}
	return
}

func (t *tabunganFeature) HistoryFeature(ctx context.Context, noRekening string) (response []model.MutasiJSON, err error) {
	getData, err := t.tabunganRepository.GetHistoryRepository(ctx, helper.StrToInt64(noRekening))
	if err != nil {
		return
	}

	if getData == nil {
		err = Error.New(ctx, shared_constant.ErrGeneral, constant.ErrNoRekeningNotFound, errors.New(fmt.Sprintf(constant.ErrNoRekeningNotFoundWithNoRekening, noRekening)))
		return
	}

	for _, v := range getData {
		response = append(response, model.MutasiJSON{
			KodeTransaksi: v.KodeTransaksi,
			Nominal:       v.Nominal,
			Saldo:         v.Saldo,
			Waktu:         v.CreatedAt,
		})
	}
	return
}
