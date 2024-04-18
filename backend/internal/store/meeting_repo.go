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

func (pg *PgMeetingRepo) GetOneByID(id string) (*models.Meeting, error) {
    query := "SELECT * FROM meetings WHERE id = $1;"

    row := pg.Db.QueryRow(query, id)
    if row.Err() != nil {
        return nil, row.Err()
    }

    var m models.Meeting
    row.Scan(&m)

    query = "SELECT * FROM participants WHERE meeting_id = $1;"

    rows, err := pg.Db.Query(query, m.ID)
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
    return &m, nil
}

// TODO: see if there is a way to merge the two queries together
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

        pQuery := "SELECT * FROM participants WHERE meeting_id = $1;"

        pRows, err := pg.Db.Query(pQuery, m.ID)
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

// TODO: figure out where this should live
func doesContain(list []string, val string) bool {
    for _, element := range list {
        if element == val {
            return true
        }
    }

    return false
}

func (pg *PgMeetingRepo) Update(m models.Meeting) error {
    query := "SELECT * FROM participants WHERE meeting_id = $1"

    rows, err := pg.Db.Query(query, m.ID)
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
            query := "DELETE FROM participants WHERE id = $1;"
            _, err := pg.Db.Exec(query, oldParticipant.ID)
            if err != nil {
                return err
            }
        }
    }

    for _, participant := range m.Participants {
        if !doesContain(oldParticipantsID, participant) {
            participantsQuery := "INSERT INTO participants (meeting_id, connection_id, owner_id) VALUES ($1, $2, $3);"

            _, err := pg.Db.Exec(participantsQuery, m.ID, participant, m.OwnerID)
            if err != nil {
                return err
            }
        }
    }

    query = `
        UPDATE meetings
        SET when = $2
            location = $3,
            notes = $4,
            description = $5
        WHERE id = $1;
    `

    _, err = pg.Db.Exec(query, m.ID, m.When, m.Location, m.Notes, m.Description)
    return err
}

func (pg *PgMeetingRepo) DeleteByID(id string) error {
    query := "SELECT * FROM participants WHERE meeting_id = $1;"

    rows, err := pg.Db.Query(query, id)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var p models.Participant

        err := rows.Scan(&p.ID, &p.MeetingID, &p.ConnectionID, &p.OwnerID)
        if err != nil {
            return err
        }


        deleteQuery := "DELETE FROM participant WHERE id = $1;"
        _, err = pg.Db.Exec(deleteQuery, p.ID)
        if err != nil {
            return err
        }
    }

    if err := rows.Err(); err != nil {
        return err
    }

    query = "DELETE FROM meetings WHERE id = $1;"
    _, err = pg.Db.Exec(query, id)
    return err
}

