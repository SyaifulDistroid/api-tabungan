package repository

import (
	"api-tabungan/infrastructure/database"
	"context"

	"api-tabungan/domain/shared/constant"
	Error "api-tabungan/domain/shared/error"
)

type HealthRepository interface {
	DatabaseHealth(ctx context.Context) (status bool, err error)
}

type healthRepository struct {
	database *database.Database
}

func NewHealthRepository(db *database.Database) HealthRepository {
	return &healthRepository{
		database: db,
	}
}

func (hr healthRepository) DatabaseHealth(ctx context.Context) (status bool, err error) {
	if hr.database.DB != nil {
		status = true
	} else {
		return
	}

	_, err = hr.database.DB.QueryContext(ctx, "SELECT 1")
	if err != nil {
		err = Error.New(ctx, constant.ErrTimeout, "error when select 1 to db", err)
		return
	}

	return
}
