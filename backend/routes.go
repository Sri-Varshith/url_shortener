package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", HomeHandler)

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "URL Shortener API")
}
