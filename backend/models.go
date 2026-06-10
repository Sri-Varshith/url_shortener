package main

import (
	"time"
)

type URLData struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessCount int       `json:"accessCount"`
}

type URLRequest struct {
	URL string `json:"url"`
}
