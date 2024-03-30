package validators

import (
	"database/sql"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/Anand-S23/capsule/internal/models"
	"github.com/Anand-S23/capsule/internal/store"
)

const (
    MAX_EMAIL_LENGHT = 255
    MAX_NAME_LENGHT = 255
    MAX_PHONE_LENGHT = 20
)

func AuthValidator(userData models.RegisterDto, store *store.Store) map[string]string {
    errs := make(map[string]string, 4)

    err := validateName(userData.FirstName, userData.LastName)
    if err != nil {
        errs["name"] = err.Error()
    }

    err = validateEmail(userData.Email, store)
    if err != nil {
        errs["email"] = err.Error()
    }

    err = validatePhoneNumber(userData.Phone)
    if err != nil {
        errs["phone"] = err.Error()
    }

    err = validatePassword(userData.Password, userData.Confirm)
    if err != nil {
        errs["password"] = err.Error()
    }

    return errs
}

func validateName(firstName string, lastName string) error {
    fullName := fmt.Sprintf("%s %s", firstName, lastName)
    if len(fullName) > MAX_NAME_LENGHT {
        return errors.New("Name entered is not valid too long")
    }

    return nil
}

func validateEmail(email string, store *store.Store) error {
    _, err := mail.ParseAddress(email)
    if err != nil || len(email) > MAX_EMAIL_LENGHT {
        return errors.New("Email entered is not valid")
    }

    user, err := store.UserRepo.GetByEmail(email)
    if err != nil && err != sql.ErrNoRows {
        return errors.New("Internal server error, please try again later")
    } else if user != nil && user.ID != "" && err != sql.ErrNoRows {
        return errors.New("User already exsits with that email")
    }

    return nil
}

func validatePhoneNumber(phoneNumber string) error {
    // https://www.twilio.com/en-us/blog/validate-e164-phone-number-in-go
    e164Regex := `^\+[1-9]\d{1,14}$`
    re := regexp.MustCompile(e164Regex)
    phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

    if len(phoneNumber) > MAX_PHONE_LENGHT || re.Find([]byte(phoneNumber)) == nil {
        return errors.New("Phone number is not valid")
    }

    return nil
}

func validatePassword(password string, confirm string) error {
    if len(password) < 8 || len(password) > 30 {
        return errors.New("Password must be between 8 and 30 characters long")
    }

    if password != confirm {
        return errors.New("Passwords must match")
    }

    return nil
}

