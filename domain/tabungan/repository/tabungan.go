package repository

import (
	"api-tabungan/infrastructure/database"
)

type TabunganRepository interface {
}

type tabunganRepository struct {
	database *database.Database
}

func NewTabunganRepository(db *database.Database) TabunganRepository {
	return &tabunganRepository{
		database: db,
	}
}
