package store

import (
	"context"
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type ReminderRepo interface {
    Add(ctx context.Context, r models.Reminder) error
    DeleteByID(ctx context.Context, id string) error
    GetByID(ctx context.Context, id string) error
    GetAllByOwnerID(ctx context.Context, owner_id string) ([]*models.Reminder, error)
    GetAllByConnectionID(ctx context.Context, owner_id string, connection_id string) ([]*models.Reminder, error)
    Update(ctx context.Context, r models.Reminder) error
}

// Postgres Meeting Repo

type PgReminderRepo struct {
    Db *sql.DB
}

func NewPgReminderRepo(db *sql.DB) *PgReminderRepo {
    return &PgReminderRepo {
        Db: db,
    }
}

func (pg *PgReminderRepo) Add(ctx context.Context, r models.Reminder) error {
    statement, err := pg.Db.PrepareContext(ctx, "INSERT INTO reminders VALUES ($1, $2, $3, $4, $5, $6)")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.ExecContext(ctx, r.ID, r.ConnectionID, r.OwnerID, r.When, r.Description, r.CreatedAt)
    return err
}

func (pg *PgReminderRepo) DeleteByID(ctx context.Context, id string) error {
    statement, err := pg.Db.PrepareContext(ctx, "DELETE FROM reminders WHERE id = $1")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.ExecContext(ctx, id)
    return err
}

func (pg *PgReminderRepo) GetByID(ctx context.Context, id string) error {
    statement, err := pg.Db.PrepareContext(ctx, "SELECT id, connection_id, owner_id, time, description, created_at FROM reminders WHERE id = $1")
    if err != nil {
        return err
    }
    defer statement.Close()

    var r models.Reminder
    err = statement.QueryRowContext(ctx, id).Scan(&r.ID, &r.ConnectionID, &r.OwnerID, &r.When, &r.Description, &r.CreatedAt)
    return err
}

func (pg *PgReminderRepo) GetAllByOwnerID(ctx context.Context, owner_id string) ([]*models.Reminder, error) {
    statement, err := pg.Db.PrepareContext(ctx, "SELECT id, connection_id, owner_id, time, description, created_at FROM reminders WHERE owner_id = $1")
    if err != nil {
        return nil, err
    }
    defer statement.Close()

    rows, err := statement.QueryContext(ctx, owner_id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    reminders := []*models.Reminder{}
    for rows.Next() {
        var r models.Reminder
        err := rows.Scan(&r.ID, &r.ConnectionID, &r.OwnerID, &r.When, &r.Description, &r.CreatedAt)
        if err != nil {
            return nil, err
        }

        reminders = append(reminders, &r)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return reminders, nil
}

func (pg *PgReminderRepo) GetAllByConnectionID(ctx context.Context, owner_id string, connection_id string) ([]*models.Reminder, error) {
    statement, err := pg.Db.PrepareContext(ctx, "SELECT id, connection_id, owner_id, time, description, created_at FROM reminders WHERE owner_id = $1 and connection_id = $2")
    if err != nil {
        return nil, err
    }
    defer statement.Close()

    rows, err := statement.QueryContext(ctx, owner_id, connection_id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    reminders := []*models.Reminder{}
    for rows.Next() {
        var r models.Reminder
        err := rows.Scan(&r.ID, &r.ConnectionID, &r.OwnerID, &r.When, &r.Description, &r.CreatedAt)
        if err != nil {
            return nil, err
        }

        reminders = append(reminders, &r)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return reminders, nil
}

func (pg *PgReminderRepo) Update(ctx context.Context, r models.Reminder) error {
    statement, err := pg.Db.PrepareContext(ctx, "UPDATE reminders SET time = $2, description = $3 WHERE id = $1")
    if err != nil {
        return err
    }
    defer statement.Close()

    _, err = statement.ExecContext(ctx, r.ID, r.When, r.Description)
    return err
}

