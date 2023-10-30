package rest

import (
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}

func NewUnauthorized(message string) *Error {
	return &Error{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "Unauthorized",
	}
}

func LogError(err *Error) {
	log.Printf("status:%v \nerror:%v \nmessage:%v", err.Status, err.Error, err.Message)
}
