package auth

import (
	"errors"
	"net/mail"
	"strings"
)

func ValidateRegisterInput(input RegisterInput) error {

	if strings.TrimSpace(input.Username) == "" {
		return errors.New("username is required")
	}

	if len(input.Username) < 3 {
		return errors.New("username too short")
	}

	if strings.TrimSpace(input.Email) == "" {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if len(input.Password) < 8 {
		return errors.New("password too short")
	}

	return nil
}
