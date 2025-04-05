package model

import (
	"errors"
)

var (
	ErrDuplicateEmail    = errors.New("duplicate email")
	ErrDuplicateUsername = errors.New("duplicate username")
	ErrNotFound          = errors.New("record not found")
	ErrUnAuthorized      = errors.New("unauthorized")
	ErrInternal          = errors.New("internal error")
)

// ErrorM is used to create the validation error response format according to the API spec
type ErrorM map[string][]string

// Error is needed to implement the error interface
func (e ErrorM) Error() string {
	return "validation error"
}
