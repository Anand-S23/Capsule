package store

import (
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type ConnectionRepo interface {
    Add(models.Connection) error
    GetOneByID(id string) (*models.Connection, error)
    GetAllByOwnerID(ownerID string) (*models.Connection, error)
    Update(c models.Connection) error
    DeleteByID(id string) error
}

// Postgres Connection Repo

type PgConnectionRepo struct {
    Db *sql.DB
}

func NewPgConnectionRepo(db *sql.DB) *PgConnectionRepo {
    return &PgConnectionRepo {
        Db: db,
    }
}

func (pg *PgConnectionRepo) Add(c models.Connection) error {
    stmt, err := pg.Db.Prepare(`
        INSERT INTO users (id, owner_id, first_name, last_name, email, phone, linkedin, company, dob, notes, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
    `)

    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(
        c.ID, 
        c.OwnerID, 
        c.FirstName, 
        c.LastName, 
        c.Email, 
        c.Phone, 
        c.LinkedIn, 
        c.Company, 
        c.Dob, 
        c.Notes, 
        c.CreatedAt)

    return err
}

func (pg *PgConnectionRepo) GetOneByID(id string) (*models.Connection, error) {
    stmt, err := pg.Db.Prepare(`
        SELECT id, owner_id, first_name, last_name, email, phone, linkedin, company, dob, notes, created_at FROM connections WHERE id = $1;
    `)
	if err != nil {
        return nil, err
	}
	defer stmt.Close()

	var c models.Connection
	err = stmt.QueryRow(id).Scan(
        &c.ID, &c.OwnerID, &c.Email, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.LinkedIn, &c.Company, &c.Dob, &c.Notes, &c.CreatedAt)
	if err != nil {
        return nil, err
	}

    return &c, nil
}

func (pg *PgConnectionRepo) GetAllByOwnerID(ownerID string) (*models.Connection, error) {
    stmt, err := pg.Db.Prepare(`
        SELECT id, owner_id, first_name, last_name, email, phone, linkedin, company, dob, notes, created_at FROM connections WHERE owner_id = $1;
    `)
	if err != nil {
        return nil, err
	}
	defer stmt.Close()

	var c models.Connection
	err = stmt.QueryRow(ownerID).Scan(
        &c.ID, &c.OwnerID, &c.Email, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.LinkedIn, &c.Company, &c.Dob, &c.Notes, &c.CreatedAt)
	if err != nil {
        return nil, err
	}

    return &c, nil
}

func (pg *PgConnectionRepo) Update(c models.Connection) error {
    stmt, err := pg.Db.Prepare(`
        UPDATE connections
        SET first_name = $1
            last_name = $2,
            email = $3,
            phone = $4,
            linkedin = $5,
            company = $6,
            dob = $7
        WHERE id = $8;
    `)
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(c.FirstName, c.LastName, c.Email, c.Phone, c.LinkedIn, c.Company, c.Dob)

    return err
}

func (pg *PgConnectionRepo) DeleteByID(id string) error {
    stmt, err := pg.Db.Prepare(`
        DELETE FROM connections WHERE id = $1;
    `)
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.Exec(id)
    return err
}

