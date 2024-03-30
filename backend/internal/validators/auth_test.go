package validators

import "testing"

func TestValidateName(t *testing.T) {
    t.Run("Valid Name", func(t *testing.T) {
        firstName := "Test"
        lastName := "User"

        err := ValidateName(firstName, lastName)
        if err != nil {
            t.Error("Returned following error for valid name :: ", err)
        }
    })

    t.Run("Invalid Name: Contains Number", func(t *testing.T) {
        firstName := "Test1"
        lastName := "User"

        err := ValidateName(firstName, lastName)
        if err != ErrorNameNonAlpha  {
            t.Error("Should have been ErrorNameNonAlpha but got ", err)
        }
    })

    t.Run("Invalid Name: Too Long", func(t *testing.T) {
        firstName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
        lastName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

        err := ValidateName(firstName, lastName)
        if err != ErrorNameTooLong  {
            t.Error("Should have been ErrorNameTooLong but got ", err)
        }
    })
}

// TODO: TestValidateEmail

func TestValidatePhoneNumber(t *testing.T) {
    t.Run("Valid Phone Number", func(t *testing.T) {
        phone := "+11234567890"

        err := ValidatePhoneNumber(phone)
        if err != nil {
            t.Error("Should have been nil but got ", err)
        }
    })

    t.Run("Invalid Phone Number", func(t *testing.T) {
        phone := "4567890"

        err := ValidatePhoneNumber(phone)
        if err != ErrorPhoneNumberInvalid {
            t.Error("Should have been nil but got ", err)
        }
    })
}

func TestValidatePassword(t *testing.T) {
    t.Run("Valid Password", func(t *testing.T) {
        password := "Password123"
        confirm := "Password123"

        err := ValidatePassword(password, confirm)
        if err != nil {
            t.Error("Should have been nil but got ", err)
        }
    })

    t.Run("Invalid Password Length", func(t *testing.T) {
        password := "test"
        confirm := "test"

        err := ValidatePassword(password, confirm)
        if err != ErrorPasswordIncorrectLenght {
            t.Error("Should have been ErrorPasswordIncorrectLenght but got ", err)
        }
    })

    t.Run("Password do not match", func(t *testing.T) {
        password := "Password123"
        confirm := "Passowrd234"

        err := ValidatePassword(password, confirm)
        if err != ErrorPasswordMismatch {
            t.Error("Should have been ErrorPasswordMismatch but got ", err)
        }
    })
}

