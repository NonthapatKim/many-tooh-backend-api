package repository

import (
	"database/sql"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/port"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) port.Repository {
	return &repository{db: db}
}
