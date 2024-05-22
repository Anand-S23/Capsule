package validators

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/Anand-S23/capsule/internal/models"
	"github.com/Anand-S23/capsule/internal/store"
)

const (
    MAX_EMAIL_LENGHT = 255
    MAX_NAME_LENGHT = 255
    MAX_PHONE_LENGHT = 20
    MIN_PASSWORD_LENGHT = 8
    MAX_PASSWORD_LENGTH = 30
)

var (
    ErrorNameNonAlpha error = errors.New("name must only contain letter")
    ErrorNameTooLong error = errors.New("name entered is not valid too long")
    ErrorEmailInvalid error = errors.New("email entered is not valid")
    ErrorEmailExists error = errors.New("user already exsits with that email")
    ErrorPhoneNumberInvalid error = errors.New("phone number is not valid")
    ErrorPasswordMismatch error = errors.New("passwords do not match")
    ErrorPasswordIncorrectLenght error = errors.New(fmt.Sprintf("password must be between %d and %d characters long", MIN_PASSWORD_LENGHT, MAX_PASSWORD_LENGTH))
)

func AuthValidator(userData models.RegisterDto, store *store.Store) map[string]string {
    errs := make(map[string]string, 4)

    err := ValidateName(userData.FirstName, userData.LastName)
    if err != nil {
        errs["name"] = err.Error()
    }

    err = ValidateEmail(userData.Email, store)
    if err != nil {
        errs["email"] = err.Error()
    }

    err = ValidatePhoneNumber(userData.Phone)
    if err != nil {
        errs["phone"] = err.Error()
    }

    err = ValidatePassword(userData.Password, userData.Confirm)
    if err != nil {
        errs["password"] = err.Error()
    }

    return errs
}

func ValidateName(firstName string, lastName string) error {
     for _, l := range firstName {
        if !unicode.IsLetter(l) {
            return ErrorNameNonAlpha
        }
    }

     for _, l := range lastName {
        if !unicode.IsLetter(l) {
            return ErrorNameNonAlpha
        }
    }

    fullName := fmt.Sprintf("%s %s", firstName, lastName)
    if len(fullName) > MAX_NAME_LENGHT {
        return ErrorNameTooLong
    }

    return nil
}

func ValidateEmail(email string, store *store.Store) error {
    _, err := mail.ParseAddress(email)
    if err != nil || len(email) > MAX_EMAIL_LENGHT {
        return ErrorEmailInvalid
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    user, err := store.UserRepo.GetByEmail(ctx, email)
    if err != nil && err != sql.ErrNoRows {
        return errors.New("internal server error, please try again later")
    } else if user != nil && user.ID != "" && err != sql.ErrNoRows {
        return ErrorEmailExists
    }

    return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
    // https://www.twilio.com/en-us/blog/validate-e164-phone-number-in-go
    e164Regex := `^\+[1-9]\d{1,14}$`
    re := regexp.MustCompile(e164Regex)
    phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

    if len(phoneNumber) > MAX_PHONE_LENGHT || re.Find([]byte(phoneNumber)) == nil {
        return ErrorPhoneNumberInvalid
    }

    return nil
}

func ValidatePassword(password string, confirm string) error {
    if len(password) < MIN_PASSWORD_LENGHT || len(password) > MAX_PASSWORD_LENGTH {
        return ErrorPasswordIncorrectLenght
    }

    if password != confirm {
        return ErrorPasswordMismatch
    }

    return nil
}

