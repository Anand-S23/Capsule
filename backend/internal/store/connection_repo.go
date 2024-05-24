package store

import (
	"context"
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type ConnectionRepo interface {
    Add(context.Context, models.Connection) error
    GetOneByID(context.Context, string) (*models.Connection, error)
    GetAllByOwnerID(context.Context, string) ([]*models.Connection, error)
    Update(context.Context, models.Connection) error
    DeleteByID(context.Context, string) error
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

func (pg *PgConnectionRepo) Add(ctx context.Context, c models.Connection) error {
    statement, err := pg.Db.PrepareContext(
        ctx,
        "INSERT INTO connections VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
    )
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.ExecContext(ctx, c.ID, c.OwnerID, c.FirstName, c.LastName, c.Email, c.Phone, 
        c.LinkedIn, c.Company, c.Dob, c.Notes, c.CreatedAt)
    return err
}

func (pg *PgConnectionRepo) GetOneByID(ctx context.Context, id string) (*models.Connection, error) {
    statement, err := pg.Db.PrepareContext(
        ctx, 
        "SELECT id, owner_id, first_name, last_name, email, phone, linkedin, company, dob, notes, created_at FROM connections WHERE id = $1;",
    )
    if err != nil {
        return nil, err
    }
    defer statement.Close()

    row := statement.QueryRowContext(ctx, id)
    if row.Err() != nil {
        return nil, row.Err()
    }

    var c models.Connection
    row.Scan(&c.ID, &c.OwnerID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.LinkedIn, &c.Company, &c.Dob, &c.Notes, &c.CreatedAt)

    return &c, nil
}

func (pg *PgConnectionRepo) GetAllByOwnerID(ctx context.Context, ownerID string) ([]*models.Connection, error) {
    statement, err := pg.Db.PrepareContext(ctx, "SELECT * FROM connections WHERE owner_id = $1;")
    if err != nil {
        return nil, err
    }
    defer statement.Close()

    rows, err := statement.QueryContext(ctx, ownerID)
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

func (pg *PgConnectionRepo) Update(ctx context.Context, c models.Connection) error {
    statement, err := pg.Db.PrepareContext(
        ctx, 
        `UPDATE connections
        SET first_name = $2
            last_name = $3,
            email = $4,
            phone = $5,
            linkedin = $6,
            company = $7,
            dob = $8
        WHERE id = $1;`,
    )
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.ExecContext(ctx, c.ID, c.FirstName, c.LastName, 
        c.Email, c.Phone, c.LinkedIn, c.Company, c.Dob) 
    return err
}

func (pg *PgConnectionRepo) DeleteByID(ctx context.Context, id string) error {
    statement, err := pg.Db.PrepareContext(ctx, "DELETE FROM connections WHERE id = $1;")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.Exec(ctx, id)
    return err
}

