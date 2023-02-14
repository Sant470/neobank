package api

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	v1 "github.com/sant470/neobank/apis/v1"
)

func InitAccountRoutes(r chi.Router, ah *v1.AccountHandler) {
	r.Route("/api/v1/account", func(r chi.Router){
		r.Method(http.MethodPost, "/", Handler(ah.CreateAccount))
	})
}