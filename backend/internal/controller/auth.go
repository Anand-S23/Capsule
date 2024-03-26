package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/capsule/internal/jwt"
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
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error occured, please try again later"))
	}
    userData.Password = string(hashedPassword)

    user := models.NewUser(userData)
    err = c.store.UserRepo.Add(user)
    if err != nil {
        log.Printf("Error storing the password in the database, %s\n", err)
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error occured, please try again later"))
    }

    successMsg := map[string]string {
        "message": "User created successfully",
        "userID": user.ID,
    }

    log.Println("User created successfully")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) error {
    var loginData models.LoginDto
    err := json.NewDecoder(r.Body).Decode(&loginData)
    if err != nil {
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Could not parse login data"))
    }

    user, err := c.store.UserRepo.GetByEmail(loginData.Email)
    if user.ID == "" || err != nil {
        log.Println("Could not get user by email from database, possibly does not exist")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Incorrect email or password, please try again"))
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
        log.Println("Passwords do not match")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Incorrect email or password, please try again"))
	}

    expDuration := time.Hour * 24
    token, err := jwt.GenerateToken(c.JwtSecretKey, user.ID, expDuration)
    if err != nil {
        log.Println("Error generating token")
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error occured, please try again later"))
    }

    cookie := jwt.GenerateCookie(c.CookieSecret, jwt.COOKIE_NAME, token, expDuration, c.production)
    if cookie == nil {
        log.Println("Error generating cookie")
        return WriteJSON(w, http.StatusInternalServerError, ErrMsg("Internal server error occured, please try again later"))
    }
    http.SetCookie(w, cookie)

    successMsg := map[string]string {
        "message": "User logged in successfully",
    }
    log.Println("User successfully logged in")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) error {
    cookie := models.GenerateExpiredCookie(models.COOKIE_NAME)
    http.SetCookie(w, cookie)
    log.Println("User successfully logged out")
    return WriteJSON(w, http.StatusOK, "")
}

