package utils

import (
	"math/rand"
	"time"
)

const codeBytes = "0123456789"

// GenerateRandomCode generates a random code consisting of digits only
func GenerateRandomCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = codeBytes[rand.Intn(len(codeBytes))]
	}
	return string(b)
}

// GetRandomElement returns a random element from a list
func GetRandomElement(list interface{}) string {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(len(list.([]string)))
	return list.([]string)[randomIndex]
}