package store

import (
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type UserRepo interface {
    Add(user models.User) error
    DeleteByID(id string) error
    GetByID(id string) (*models.User, error)
    GetByEmail(email string) (*models.User, error)
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

func (pg *PgUserRepo) Add(user models.User) error {
    stmt, err := pg.Db.Prepare(`
        INSERT INTO users (id, owner_id, first_name, last_name, email, phone, linkedin, company, dob, notes, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
    `)
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(user.ID, user.Name, user.Email, user.Phone, user.Password, user.CreatedAt)
    return err
}

func (pg *PgUserRepo) DeleteByID(id string) error {
    stmt, err := pg.Db.Prepare(`
        DELETE FROM users WHERE id = $1;
    `)
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(id)
    return err
}

func (pg *PgUserRepo) GetByID(id string) (*models.User, error) {
    stmt, err := pg.Db.Prepare(`
        SELECT id, name, email, phone, password, created_at FROM users WHERE id = $1;
    `)
	if err != nil {
        return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.CreatedAt)
	if err != nil {
        return nil, err
	}

    return &user, nil
}

func (pg *PgUserRepo) GetByEmail(email string) (*models.User, error) {
    stmt, err := pg.Db.Prepare(`
        SELECT id, name, email, phone, password, created_at FROM users WHERE email = $1;
    `)
	if err != nil {
        return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.CreatedAt)
	if err != nil {
        return nil, err
	}

    return &user, nil
}

// MockUserRepo

type MockUserRepo struct {}

func (mr *MockUserRepo) Add(user models.User) error {
    return nil
}

func (mr *MockUserRepo) DeleteByID(id string) error {
    return nil
}

func (mr *MockUserRepo) GetByID(id string) (*models.User, error) {
    return nil, nil
}

func (mr *MockUserRepo) GetByEmail(email string) (*models.User, error) {
    return nil, nil
}
