package user

import (
	"errors"
	"strings"
)

var ErrEmptyEmail error = errors.New("email is required")
var ErrInvalidEmail error = errors.New("email format is invalid")

type Email struct {
	email string
}

func NewEmailVo(email string) *Email {
	return &Email{email}
}

func (e *Email) ValidateRequired() error {
	if e.email == "" {
		return ErrEmptyEmail
	}
	return nil
}

func (e *Email) ValidateFormat() error {
	if e.email != "" && !strings.Contains(e.email, "@") {
		return ErrInvalidEmail
	}
	return nil
}

func (e *Email) Validate() error {
	if err := e.ValidateRequired(); err != nil {
		return err
	}
	if err := e.ValidateFormat(); err != nil {
		return err
	}
	return nil
}

func (e *Email) Value() string {
	return e.email
}
