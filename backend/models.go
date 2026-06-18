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

type ShortCodeResponse struct {
	ShortCode string `json:"shortCode"`
}

type StatsResponse struct {
	URL         string `json:"url"`
	ShortCode   string `json:"shortCode"`
	AccessCount int    `json:"accessCount"`
}
