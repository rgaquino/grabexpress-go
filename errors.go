package grabexpress

import (
	"errors"
	"net/http"
)

var (
	ErrCredentialsMissing  = errors.New("API Key credentials missing")
	ErrAuthenticationError = errors.New("authentication error")
	ErrBaseURLMissing      = errors.New("base URL missing")
	ErrTokenURLMissing     = errors.New("token URL missing")
)

// Error is the conventional GrabExpress client error
type Error struct {
	Status    int    `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestID string `json:"requestID,omitempty"`
}

// Error returns error message.
// This enables grabexpress.Error to comply with Go error interface
func (e *Error) Error() string {
	return e.Message
}

func wrapError(err error) *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}
