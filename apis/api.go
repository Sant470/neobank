package api

import (
	"net/http"

)

type Handler func(rw http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := h(rw, r); err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("Internal server error"))
	}
}



