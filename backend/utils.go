package main

import (
	"crypto/rand"
	"math/big"
)

func GenerateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := ""

	for i := 0; i < 6; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}

		result += string(charset[randomIndex.Int64()])
	}

	return result
}
