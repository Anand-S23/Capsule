package store

import (
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

func (pg *PgMeetingRepo) Add(m models.Meeting) error {
    meetingQuery := "INSERT INTO meetings VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

    _, err := pg.Db.Exec(meetingQuery, m.ID, m.OwnerID, m.When, m.Location, m.MeetingType, 
        m.Notes, m.Description, m.CreatedAt)

    for _, participant := range m.Participants {
        participantsQuery := "INSERT INTO participants (meeting_id, connection_id, owner_id) VALUES ($1, $2, $3);"

        _, err := pg.Db.Exec(participantsQuery, m.ID, participant, m.OwnerID)
        if err != nil {
            return err
        }
    }

    return err
}

// TODO: Get Participants for the meeting
func (pg *PgMeetingRepo) GetOneByID(id string) (*models.Meeting, error) {
    query := "SELECT * FROM meetings WHERE id = $1;"

    row := pg.Db.QueryRow(query, id)
    if row.Err() != nil {
        return nil, row.Err()
    }

    var m models.Meeting
    row.Scan(&m)

    return &m, nil
}

// TODO: Get Participants for each all meetings
func (pg *PgMeetingRepo) GetAllByOwnerID(ownerID string) ([]*models.Meeting, error) {
    query := "SELECT * FROM meetings WHERE owner_id = $1;"

    rows, err := pg.Db.Query(query, ownerID)
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

        meetings = append(meetings, &m)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return meetings, nil
}

// TODO: Get Participants, compare new update list and make neccessary changes - deleting or adding new participant
func (pg *PgMeetingRepo) Update(m models.Meeting) error {
    query := `
        UPDATE meetings
        SET when = $2
            location = $3,
            notes = $4,
            description = $5
        WHERE id = $1;
    `

    _, err := pg.Db.Exec(query, m.ID, m.When, m.Location, m.Notes, m.Description)

    return err
}

// TODO: Delete Participants with the same id
func (pg *PgMeetingRepo) DeleteByID(id string) error {
    query := "DELETE FROM meetings WHERE id = $1;"

    _, err := pg.Db.Exec(query, id)
    return err
}

