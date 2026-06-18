package main

import (
	"net/http"
	"encoding/json"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var url URLRequest
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	shortCode := GenerateCode()
	

}

func GetShortURL(w http.ResponseWriter, r *http.Request) {

}

func UpdateShortURL(w http.ResponseWriter, r *http.Request) {

}

func DeleteShortURL(w http.ResponseWriter, r *http.Request) {

}

func GetStats(w http.ResponseWriter, r *http.Request) {

}
