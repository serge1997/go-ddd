package domain

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("not found")
var ErrAlwreadyExists = errors.New("already exists")
var ErrValidation = errors.New("invalid data")

func StatusFromError(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(ErrAlwreadyExists, err) || errors.Is(ErrValidation, err):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
