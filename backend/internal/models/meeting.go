package models

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
    ID           int
    MeetingID    string
    ConnectionID string
    OwnerID      string
}

type MeetingDto struct {
    When         time.Time
    Location     string
    Participants []string
    MeetingType  MeetingType
    Notes        string
    Description  string
}

type MeetingType struct {
    ID          string `json:"id"`
    OwnerID     string `json:"ownerId"`
    Description string `json:"description"`
}

type Meeting struct {
    ID           string      `json:"id"`
    OwnerID      string      `json:"ownerId"`
    When         time.Time   `json:"when"`
    Location     string      `json:"location"`
    Participants []string    `json:"participants"`
    MeetingType  MeetingType `json:"meetingType"`
    Notes        string      `json:"notes"`
    Description  string      `json:"description"`
    CreatedAt    time.Time   `json:"createdAt"`
}

func NewMeetingType(ownerId string, desc string) MeetingType {
    return MeetingType {
        ID: uuid.NewString(),
        OwnerID: ownerId,
        Description: desc,
    }
}

func NewMeeting(ownerId string, meetingDto MeetingDto) Meeting {
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return Meeting {
        ID: uuid.NewString(),
        OwnerID: ownerId,
        When: meetingDto.When,
        Location: meetingDto.Location,
        Participants: meetingDto.Participants,
        MeetingType: meetingDto.MeetingType,
        Notes: meetingDto.Notes,
        Description: meetingDto.Description,
        CreatedAt: now,
    }
}

