package apperrors

import (
	"net/http"
)

type AppError struct {
	Message    string `json:"message"`
	Code       string `json:"code,omitempty"`
	StatusCode int    `json:"statusCode"`
}

func (e *AppError) Error() string { return e.Message }

var (
	ErrInternalServerError = AppError{
		Message:    "Internal server error",
		Code:       "internal_error",
		StatusCode: http.StatusInternalServerError}
)
