package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID        string    `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Phone     string    `json:"phone"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"createdAt"`
}

func NewUser(userData RegisterDto) User {
    id := uuid.New().String()
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return User {
        ID: id,
        Email: userData.Email,
        Name: fmt.Sprintf("%s %s", userData.FirstName, userData.LastName),
        Phone: userData.Phone,
        Password: userData.Password,
        CreatedAt: now,
    }
}

