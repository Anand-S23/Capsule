package store

import (
	"context"
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type UserRepo interface {
    Add(context.Context, models.User) error
    DeleteByID(context.Context, string) error
    GetByID(context.Context, string) (*models.User, error)
    GetByEmail(context.Context, string) (*models.User, error)
}

// PostgresUserRepo

type PgUserRepo struct {
    Db *sql.DB
}

func NewPgUserRepo(db *sql.DB) *PgUserRepo{
    return &PgUserRepo {
        Db: db,
    }
}

func (pg *PgUserRepo) Add(ctx context.Context, user models.User) error {
    stmt, err := pg.Db.PrepareContext(
        ctx, 
        `INSERT INTO users (id, owner_id, first_name, last_name, email, phone, linkedin, company, dob, notes, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`,
    )
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Email, user.Phone, user.Password, user.CreatedAt)
    return err
}

func (pg *PgUserRepo) DeleteByID(ctx context.Context, id string) error {
    stmt, err := pg.Db.PrepareContext(ctx, "DELETE FROM users WHERE id = $1;")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.ExecContext(ctx, id)
    return err
}

func (pg *PgUserRepo) GetByID(ctx context.Context, id string) (*models.User, error) {
    stmt, err := pg.Db.PrepareContext(
        ctx,
        "SELECT id, name, email, phone, password, created_at FROM users WHERE id = $1;",
    )
	if err != nil {
        return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRowContext(ctx, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.CreatedAt)
	if err != nil {
        return nil, err
	}

    return &user, nil
}

func (pg *PgUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    stmt, err := pg.Db.PrepareContext(
        ctx,
        "SELECT id, name, email, phone, password, created_at FROM users WHERE email = $1;",
    )
	if err != nil {
        return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRowContext(ctx, email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.CreatedAt)
	if err != nil {
        return nil, err
	}

    return &user, nil
}

// MockUserRepo

type MockUserRepo struct {}

func (mr *MockUserRepo) Add(ctx context.Context, user models.User) error {
    return nil
}

func (mr *MockUserRepo) DeleteByID(ctx context.Context, id string) error {
    return nil
}

func (mr *MockUserRepo) GetByID(ctx context.Context, id string) (*models.User, error) {
    return nil, nil
}

func (mr *MockUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    return nil, nil
}

