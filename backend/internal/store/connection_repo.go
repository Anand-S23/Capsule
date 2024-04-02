package store

import (
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type ConnectionRepo interface {
    Add(user models.User) error
    DeleteByID(id string) error
    GetByID(id string) (*models.User, error)
    GetByEmail(email string) (*models.User, error)
}

// Postgres Connection Repo

type PgConnectionRepo struct {
    Db *sql.DB
}

