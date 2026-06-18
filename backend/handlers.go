package main

import (
	"encoding/json"
	"net/http"
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

func GetShortURL(w http.ResponseWriter, r *http.Request) {

}

func UpdateShortURL(w http.ResponseWriter, r *http.Request) {

}

func DeleteShortURL(w http.ResponseWriter, r *http.Request) {

}

func GetStats(w http.ResponseWriter, r *http.Request) {

}
