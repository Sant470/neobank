package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int 
	Message string
}

func BadRequest(message string) *AppError {
	return &AppError{StatusCode: http.StatusBadRequest, Message: message}
}

func InternalServerErrorStd() *AppError {
	return &AppError{StatusCode: http.StatusInternalServerError, Message: "something went wrong"}
}

func InternalServerError(message string) *AppError {
	return &AppError{StatusCode: http.StatusInternalServerError, Message: message}
}

func NotFound(message string) *AppError {
	return &AppError{StatusCode: http.StatusNotFound, Message: message}
}

func (err *AppError) Debug(message string) {
	fmt.Printf("error %v", message)
}