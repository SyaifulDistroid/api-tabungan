package repository

import (
	"api-tabungan/domain/tabungan/constant"
	"api-tabungan/domain/tabungan/model"
	"api-tabungan/infrastructure/database"
	"context"
	"database/sql"
	"fmt"
)

type TabunganRepository interface {
	CreateRekeningRepository(ctx context.Context, rek *model.Rekening) (err error)
	GetByNIK(ctx context.Context, nik string) (rekening *model.Rekening, err error)
	GetByNoHandphone(ctx context.Context, hp string) (rekening *model.Rekening, err error)
	GetByNoRekening(ctx context.Context, norek int64) (rekening *model.Rekening, err error)
	UpdateSaldoRepository(ctx context.Context, rek *model.Rekening) (err error)
	CreateHistoryRepository(ctx context.Context, rek *model.Mutasi) (err error)
	GetHistoryRepository(ctx context.Context, norek int64) (rekening []*model.Mutasi, err error)
}

type tabunganRepository struct {
	database *database.Database
}

func NewTabunganRepository(db *database.Database) TabunganRepository {
	return &tabunganRepository{
		database: db,
	}
}

func (t *tabunganRepository) CreateRekeningRepository(ctx context.Context, rek *model.Rekening) (err error) {
	_, err = t.database.ExecContext(ctx, CreateRekeningQuery,
		rek.Nama,
		rek.NIK,
		rek.NoHandphone,
		rek.NoRekening,
		rek.CreatedBy,
	)
	if err != nil {
		return
	}
	return
}

func (t *tabunganRepository) GetByNIK(ctx context.Context, nik string) (rekening *model.Rekening, err error) {
	paramNIK := fmt.Sprintf(constant.PARAMETER_NIK)
	query := GetRekeningQuery + paramNIK

	var rek model.Rekening
	err = t.database.GetContext(ctx, &rek, query, nik)

	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
	}
	rekening = &rek
	return
}

func (t *tabunganRepository) GetByNoHandphone(ctx context.Context, hp string) (rekening *model.Rekening, err error) {
	paramHP := fmt.Sprintf(constant.PARAMETER_NOHANDPHONE)
	query := GetRekeningQuery + paramHP

	var rek model.Rekening
	err = t.database.GetContext(ctx, &rek, query, hp)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
	}

	rekening = &rek

	return
}

func (t *tabunganRepository) GetByNoRekening(ctx context.Context, norek int64) (rekening *model.Rekening, err error) {
	paramHP := fmt.Sprintf(constant.PARAMETER_NOREKENING)
	query := GetRekeningQuery + paramHP

	var rek model.Rekening
	err = t.database.GetContext(ctx, &rek, query, norek)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
	}

	rekening = &rek

	return
}

func (t *tabunganRepository) UpdateSaldoRepository(ctx context.Context, rek *model.Rekening) (err error) {
	_, err = t.database.ExecContext(ctx, UpdateSaldoQuery,
		rek.ID,
		rek.Saldo,
	)
	if err != nil {
		return
	}
	return
}

func (t *tabunganRepository) CreateHistoryRepository(ctx context.Context, rek *model.Mutasi) (err error) {
	_, err = t.database.ExecContext(ctx, CreateHistoryQuery,
		rek.NoRekening,
		rek.KodeTransaksi,
		rek.Nominal,
		rek.Saldo,
		rek.CreatedBy,
	)
	if err != nil {
		return
	}
	return
}

func (t *tabunganRepository) GetHistoryRepository(ctx context.Context, norek int64) (rekening []*model.Mutasi, err error) {
	err = t.database.SelectContext(ctx, &rekening, GetHistoryQuery, norek)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
	}
	return
}
