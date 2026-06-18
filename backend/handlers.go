package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var url URLRequest
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var shortCode string
	for i := 0; i < 5; i++ {
		shortCode = GenerateCode()
		err := CreateURL(url.URL, shortCode)
		if err == nil {
			break
		}
	}
	response := ShortCodeResponse{
		ShortCode: shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	url, err := GetURLByCode(code)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func UpdateShortURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	var req URLRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = UpdateURL(code, req.URL)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteShortURL(w http.ResponseWriter, r *http.Request) {

}

func GetStats(w http.ResponseWriter, r *http.Request) {

}
