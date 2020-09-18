package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomCode generates a random code consisting of digits only
func GenerateRandomCode(n int, digits string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	return string(b)
}

// GetRandomElement returns a random element from a list
func GetRandomElement(list interface{}) string {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(len(list.([]string)))
	return list.([]string)[randomIndex]
}