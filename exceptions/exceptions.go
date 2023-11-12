package exceptions

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func IsNotFoundError(err error) bool {
	if customErr, ok := err.(*CustomError); ok {
		return customErr.Code == 404
	}
	return false
}
