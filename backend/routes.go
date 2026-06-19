package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
		},
	}))

	r.Post("/shorten", CreateShortURL)

	r.Get("/{code}", RedirectURL)

	r.Put("/shorten/{code}", UpdateShortURL)

	r.Delete("/shorten/{code}", DeleteShortURL)

	r.Get("/shorten/{code}/stats", GetStats)

	return r
}
