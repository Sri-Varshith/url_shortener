package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB(databaseURL string) {
	var err error

	DB, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")
}

func CreateURL(url string, shortCode string) error {
	_, err := DB.Exec(
		context.Background(),
		`INSERT INTO urls (url, short_code)
		 VALUES ($1, $2)`,
		url,
		shortCode,
	)
	return err
}

func GetURLByCode(code string) (string, error) {
	var url string

	err := DB.QueryRow(
		context.Background(),
		`SELECT url FROM urls WHERE short_code = $1`,
		code,
	).Scan(&url)

	return url, err
}

func UpdateURL(code string, url string) error {
	_, err := DB.Exec(
		context.Background(),
		`UPDATE urls
		 SET url = $1
		 WHERE short_code = $2`,
		url,
		code,
	)

	return err
}

func DeleteURL(code string) error {
	_, err := DB.Exec(
		context.Background(),
		`DELETE FROM urls
		 WHERE short_code = $1`,
		code,
	)

	return err
}

func IncrementAccessCount(code string) error {
	_, err := DB.Exec(
		context.Background(),
		`UPDATE urls
		 SET access_count = access_count + 1
		 WHERE short_code = $1`,
		code,
	)

	return err
}

func GetURLStats(code string) (StatsResponse, error) {
	var data StatsResponse

	err := DB.QueryRow(
		context.Background(),
		`SELECT url, short_code, access_count
		 FROM urls
		 WHERE short_code = $1`,
		code,
	).Scan(
		&data.URL,
		&data.ShortCode,
		&data.AccessCount,
	)

	return data, err
}