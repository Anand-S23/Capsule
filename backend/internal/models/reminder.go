package models

import (
	"time"

	"github.com/google/uuid"
)

type ReminderDto struct {
    ConnectionID string
    When         time.Time
    Description  string
}

type Reminder struct {
    ID           string
    OwnerID      string
    ConnectionID string
    When         time.Time
    Description  string
    CreatedAt    time.Time
}

func NewReminder(ownerId string, reminderDto ReminderDto) Reminder {
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return Reminder {
        ID: uuid.NewString(),
        OwnerID: ownerId,
        ConnectionID: reminderDto.ConnectionID,
        When: reminderDto.When,
        Description: reminderDto.Description,
        CreatedAt: now,
    }
}

