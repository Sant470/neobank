package api

import (
	"net/http"
	errors "github.com/sant470/neobank/utils/errors"
)

type Handler func(rw http.ResponseWriter, r *http.Request) *errors.AppError


func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := h(rw, r); err != nil {
		rw.WriteHeader(err.StatusCode)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(err.Message))
	}
}



