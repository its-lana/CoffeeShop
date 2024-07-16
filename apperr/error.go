package apperr

import (
	"errors"
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (c *CustomError) Error() string {
	return c.Message
}

func (c *CustomError) ConvertToErrorResponse() ErrorResponse {
	return ErrorResponse{
		Message: c.Message,
	}
}

var (
	ErrWrongCredentialsLogin = errors.New("wrong email or password")
	ErrWrongCredentials      = NewCustomError(http.StatusUnauthorized, "wrong credentials")

	ErrInvalidBody        = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrUnathorized        = NewCustomError(http.StatusUnauthorized, "you are not authorized")
	ErrBearerTokenInvalid = NewCustomError(http.StatusUnauthorized, "bearer token is invalid")
)
