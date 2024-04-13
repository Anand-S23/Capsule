package models

import (
	"time"

	"github.com/google/uuid"
)

type ConnectionDto struct {
    FirstName string
    LastName  string
    Email     string
    Phone     string
    LinkedIn  string
    Company   string
    Dob       time.Time
    Notes     string
}

type Connection struct {
    ID        string    `json:"id"`
    OwnerID   string    `json:"ownerId"`
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    LinkedIn  string    `json:"linkedin"`
    Company   string    `json:"company"`
    Dob       time.Time `json:"dob"`
    Notes     string    `json:"notes"`
    CreatedAt time.Time `json:"createdAt"`
}

func NewConnection(connectionData ConnectionDto, ownerId string) Connection {
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return Connection {
        ID: uuid.NewString(),
        OwnerID: ownerId,
        FirstName: connectionData.FirstName,
        LastName: connectionData.LastName,
        Email: connectionData.Email,
        Phone: connectionData.Phone,
        LinkedIn: connectionData.LinkedIn,
        Company: connectionData.Company,
        Dob: connectionData.Dob,
        CreatedAt: now,
    }
}

func NewConnectionWithID(connectionData ConnectionDto, ownerId string, connectionID string) Connection {
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return Connection {
        ID: connectionID,
        OwnerID: ownerId,
        FirstName: connectionData.FirstName,
        LastName: connectionData.LastName,
        Email: connectionData.Email,
        Phone: connectionData.Phone,
        LinkedIn: connectionData.LinkedIn,
        Company: connectionData.Company,
        Dob: connectionData.Dob,
        CreatedAt: now,
    }
}

