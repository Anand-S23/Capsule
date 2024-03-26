package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Anand-S23/capsule/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) error {
    var userData models.RegisterDto
    err := json.NewDecoder(r.Body).Decode(&userData)
    if err != nil {
        log.Println("Error parsing register data")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Could not parse register data"))
    }

    // TODO: Input validation

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
        log.Println("Error hashing the password")
        return WriteJSON(w, http.StatusInternalServerError, "Internal server error occured, please try again later")
	}
    userData.Password = string(hashedPassword)

    user := models.NewUser(userData)
    err = c.store.UserRepo.Add(user)
    if err != nil {
        log.Printf("Error storing the password in the database, %s\n", err)
        return WriteJSON(w, http.StatusInternalServerError, "Internal server error occured, please try again later")
    }

    successMsg := map[string]string {
        "message": "User created successfully",
        "userID": user.ID,
    }

    log.Println("User created successfully")
    return WriteJSON(w, http.StatusOK, successMsg)
}

