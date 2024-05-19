package store

import (
	"context"
	"database/sql"

	"github.com/Anand-S23/capsule/internal/models"
)

type MeetingRepo interface {
    Add(models.Meeting) error
    GetOneByID(id string) (*models.Meeting, error)
    GetAllByOwnerID(ownerID string) ([]*models.Meeting, error)
    Update(c models.Meeting) error
    DeleteByID(id string) error
}

// Postgres Meeting Repo

type PgMeetingRepo struct {
    Db *sql.DB
}

func NewPgMeetingRepo(db *sql.DB) *PgMeetingRepo {
    return &PgMeetingRepo {
        Db: db,
    }
}

func (pg *PgMeetingRepo) Add(ctx context.Context, m models.Meeting) (err error) {
    tx, err := pg.Db.BeginTx(ctx, nil)
    if err != nil {
        return
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    meetingQuery, err := tx.PrepareContext(ctx, "INSERT INTO meetings VALUES ($1, $2, $3, $4, $5, $6, $7, $8);")
    if err != nil {
        return
    }
    defer meetingQuery.Close()

    _, err = meetingQuery.ExecContext(ctx, m.ID, m.OwnerID, m.When, m.Location, m.MeetingType, 
        m.Notes, m.Description, m.CreatedAt)
    if err != nil {
        return
    }

    for _, participant := range m.Participants {
        participantsQuery, err := tx.PrepareContext(ctx, "INSERT INTO participants (meeting_id, connection_id, owner_id) VALUES ($1, $2, $3);")
        if err != nil {
            return err
        }
        defer participantsQuery.Close()

        _, err = participantsQuery.ExecContext(ctx, m.ID, participant, m.OwnerID)
        if err != nil {
            return err
        }
    }

    return
}

func (pg *PgMeetingRepo) GetOneByID(ctx context.Context, id string) (m *models.Meeting, err error) {
    tx, err := pg.Db.BeginTx(ctx, nil)
    if err != nil {
        return nil, err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    statement, err := tx.PrepareContext(
        ctx, 
        "SELECT id, owner_id, time, location, description, notes, created_at, FROM meetings WHERE id = $1;",
    )
    if err != nil {
        return nil, err
    }
    defer statement.Close()

    err = statement.QueryRowContext(ctx, id).Scan(&m.ID, &m.OwnerID, &m.When, &m.Location, &m.Description, &m.Notes, &m.CreatedAt)
	if err != nil {
        return nil, err
	}
    
    pStatement, err := tx.PrepareContext(ctx, "SELECT id, meeting_id, connection_id, owner_id FROM participants WHERE meeting_id = $1;")
    if err != nil {
        return nil, err
    }

    rows, err := pStatement.QueryContext(ctx, m.ID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    participants := []string{}

    for rows.Next() {
        var p models.Participant

        err := rows.Scan(&p.ID, &p.MeetingID, &p.ConnectionID, &p.OwnerID)

        if err != nil {
            return nil, err
        }

        participants = append(participants, p.ConnectionID)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    m.Participants = participants
    return m, nil
}

func (pg *PgMeetingRepo) GetAllByOwnerID(ctx context.Context, ownerID string) ([]*models.Meeting, error) {
    tx, err := pg.Db.BeginTx(ctx, nil)
    if err != nil {
        return nil, err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    statement, err := tx.PrepareContext(
        ctx, 
        "SELECT id, owner_id, time, location, description, notes, created_at, FROM meetings WHERE owner_id = $1;",
    )
    if err != nil {
        return nil, err
    }

    rows, err := statement.QueryContext(ctx, ownerID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    meetings := []*models.Meeting{}

    for rows.Next() {
        var m models.Meeting

        err := rows.Scan(
            &m.ID, &m.OwnerID, &m.When, &m.Location, &m.Notes, &m.Description, &m.CreatedAt)

        if err != nil {
            return nil, err
        }

        pStatement, err := tx.PrepareContext(ctx, "SELECT id, meeting_id, connection_id, owner_id FROM participants WHERE meeting_id = $1;")
        if err != nil {
            return nil, err
        }

        pRows, err := pStatement.QueryContext(ctx, m.ID)
        if err != nil {
            return nil, err
        }
        defer pRows.Close()

        participants := []string{}

        for pRows.Next() {
            var p models.Participant

            err := pRows.Scan(&p.ID, &p.MeetingID, &p.ConnectionID, &p.OwnerID)

            if err != nil {
                return nil, err
            }

            participants = append(participants, p.ConnectionID)
        }

        if err := pRows.Err(); err != nil {
            return nil, err
        }

        m.Participants = participants
        meetings = append(meetings, &m)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return meetings, nil
}

func doesContain(list []string, val string) bool {
    for _, element := range list {
        if element == val {
            return true
        }
    }

    return false
}

func (pg *PgMeetingRepo) Update(ctx context.Context, m models.Meeting) error {
    tx, err := pg.Db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    pStatement, err := tx.PrepareContext(ctx, "SELECT id, meeting_id, connection_id, owner_id FROM participants WHERE meeting_id = $1;")
    if err != nil {
        return err
    }

    rows, err := pStatement.QueryContext(ctx, m.ID)
    if err != nil {
        return err
    }

    var oldParticipants []models.Participant
    var oldParticipantsID []string

    for rows.Next() {
        var p models.Participant

        err := rows.Scan(&p.ID, &p.MeetingID, &p.ConnectionID, &p.OwnerID)

        if err != nil {
            return err
        }

        oldParticipants = append(oldParticipants, p)
        oldParticipantsID = append(oldParticipantsID, p.ConnectionID)
    }

    for _, oldParticipant := range oldParticipants {
        if !doesContain(m.Participants, oldParticipant.ConnectionID) {
            pdStatement, err := tx.PrepareContext(ctx, "DELETE FROM participants WHERE id = $1;")
            if err != nil {
                return nil
            }

            _, err = pdStatement.ExecContext(ctx, oldParticipant.ID)
            if err != nil {
                return err
            }
        }
    }

    for _, participant := range m.Participants {
        if !doesContain(oldParticipantsID, participant) {
            piStatement, err := tx.PrepareContext(ctx, "INSERT INTO participants (meeting_id, connection_id, owner_id) VALUES ($1, $2, $3);")
            if err != nil {
                return nil
            }

            _, err = piStatement.ExecContext(ctx, m.ID, participant, m.OwnerID)
            if err != nil {
                return err
            }
        }
    }

    muStatement, err := tx.PrepareContext(ctx, 
        `UPDATE meetings
        SET when = $2
            location = $3,
            notes = $4,
            description = $5
        WHERE id = $1;`,
   )
   if err != nil {
       return err
   }

    _, err = muStatement.ExecContext(ctx, m.ID, m.When, m.Location, m.Notes, m.Description)
    return err
}

func (pg *PgMeetingRepo) DeleteByID(ctx context.Context, id string) error {
    tx, err := pg.Db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    psStatement, err := tx.PrepareContext(ctx, "SELECT id FROM participants WHERE meeting_id = $1;")
    if err != nil {
        return err
    }

    rows, err := psStatement.QueryContext(ctx, id)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var pid int

        err := rows.Scan(&pid)
        if err != nil {
            return err
        }

        pdStatement, err := tx.PrepareContext(ctx, "DELETE FROM participant WHERE id = $1;")
        _, err = pdStatement.ExecContext(ctx, pid)
        if err != nil {
            return err
        }
    }

    if err := rows.Err(); err != nil {
        return err
    }

    mdStatement, err := tx.PrepareContext(ctx, "DELETE FROM meetings WHERE id = $1;")
    if err != nil {
        return err
    }

    _, err = mdStatement.ExecContext(ctx, id)
    return err
}

