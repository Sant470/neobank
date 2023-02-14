package appError

import (
	"fmt"
	"net/http"
)

type appError struct {
	statusCode int 
	message string
}


func BadRequest(message string) *appError {
	return &appError{statusCode: http.StatusBadRequest, message: message}
}

func InternalServerErrorStd() *appError {
	return &appError{statusCode: http.StatusInternalServerError, message: "something went wrong"}
}

func InternalServerError(message string) *appError {
	return &appError{statusCode: http.StatusInternalServerError, message: message}
}

func NotFound(message string) *appError {
	return &appError{statusCode: http.StatusNotFound, message: message}
}

func (err *appError) Debug(message string) {
	fmt.Printf("error %v", message)
}