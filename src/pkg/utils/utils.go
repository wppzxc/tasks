package utils

import "math/rand"

func GenerateKey(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 97
		bytes[i] = byte(b)
	}
	return string(bytes)
}
