package models

import (
	"errors"
	"regexp"
)

const (
	nameIsRequiredError     = "name is required"
	emailIsRequiredError    = "email is required"
	emailRexPatter          = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailInvalidError       = "email is invalid"
	passwordIsRequiredError = "password is required"
	passwordMinLength       = 6
	passwordMinLengthError  = "password must be at least 6 characters long"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (u *User) Validate() error {

	if u.Name == "" {
		return errors.New(nameIsRequiredError)
	}

	if u.Email == "" {
		return errors.New(emailIsRequiredError)
	}

	if matched, _ := regexp.MatchString(emailRexPatter, u.Email); !matched {
		return errors.New(emailInvalidError)
	}

	if u.Password == "" {
		return errors.New(passwordIsRequiredError)
	}

	if len(u.Password) < passwordMinLength {
		return errors.New(passwordMinLengthError)
	}

	return nil
}
