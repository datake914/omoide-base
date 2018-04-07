package exception

import (
	"fmt"
	"net/http"
)

type ValidationError struct {
	messageCode    string
	message        string
	httpStatusCode int
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		messageCode:    "00400",
		message:        fmt.Sprintf("Invalid parameter value: %s", message),
		httpStatusCode: http.StatusBadRequest,
	}
}

func (err *ValidationError) Error() string {
	return fmt.Sprintf("[%s] %s", err.MessageCode, err.Message)
}

func (err *ValidationError) HttpStatusCode() int {
	return err.httpStatusCode
}

func (err *ValidationError) MessageCode() string {
	return err.messageCode
}

func (err *ValidationError) Message() string {
	return err.message
}
