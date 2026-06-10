package main

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/shorten", CreateShortURL)

	r.Get("/shorten/{code}", GetShortURL)

	r.Put("/shorten/{code}", UpdateShortURL)

	r.Delete("/shorten/{code}", DeleteShortURL)

	r.Get("/shorten/{code}/stats", GetStats)

	return r
}
