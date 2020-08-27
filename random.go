package utils

import "math/rand"

const codeBytes = "0123456789"

// GenerateRandomCode generates a random code consisting of digits only
func GenerateRandomCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = codeBytes[rand.Intn(len(codeBytes))]
	}
	return string(b)
}