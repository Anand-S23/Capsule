package store

import (
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type ConnectionRepo interface {
    Add(models.Connection) error
    GetOneByID(id string) (*models.Connection, error)
    GetAllByOwnerID(ownerID string) ([]*models.Connection, error)
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
    query := "INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"

    _, err := pg.Db.Exec(query, c.ID, c.OwnerID, c.FirstName, c.LastName, c.Email, c.Phone, 
        c.LinkedIn, c.Company, c.Dob, c.Notes, c.CreatedAt)

    return err
}

func (pg *PgConnectionRepo) GetOneByID(id string) (*models.Connection, error) {
    query := "SELECT * FROM connections WHERE id = $1;"

    row := pg.Db.QueryRow(query, id)
    if row.Err() != nil {
        return nil, row.Err()
    }

    var c models.Connection
    row.Scan(&c)

    return &c, nil
}

func (pg *PgConnectionRepo) GetAllByOwnerID(ownerID string) ([]*models.Connection, error) {
    query := "SELECT * FROM connections WHERE owner_id = $1;"

    rows, err := pg.Db.Query(query, ownerID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    connections := []*models.Connection{}

    for rows.Next() {
        var c models.Connection

        err := rows.Scan(
            &c.ID, &c.OwnerID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, 
            &c.LinkedIn, &c.Company, &c.Dob, &c.Notes, &c.CreatedAt)

        if err != nil {
            return nil, err
        }

        connections = append(connections, &c)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return connections, nil
}

func (pg *PgConnectionRepo) Update(c models.Connection) error {
    query := `
        UPDATE connections
        SET first_name = $2
            last_name = $3,
            email = $4,
            phone = $5,
            linkedin = $6,
            company = $7,
            dob = $8
        WHERE id = $1;
    `

    _, err := pg.Db.Exec(query, c.ID, c.FirstName, c.LastName, 
        c.Email, c.Phone, c.LinkedIn, c.Company, c.Dob) 

    return err
}

func (pg *PgConnectionRepo) DeleteByID(id string) error {
    query := "DELETE FROM connections WHERE id = $1;"

    _, err := pg.Db.Exec(query, id)
    return err
}

