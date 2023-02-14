package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	api "github.com/sant470/neobank/apis"
	v1 "github.com/sant470/neobank/apis/v1"
)

func main() {
	fmt.Println("project start here ...")
	lgr := log.Default()
	router := chi.NewRouter()
	ah := v1.NewAccountHandler(lgr)
	api.InitAccountRoutes(router, ah)
	lgr.Fatal(http.ListenAndServe("localhost:8000", router))
}