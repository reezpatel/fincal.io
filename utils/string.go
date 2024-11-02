package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/exp/rand"
)

func Hash(password string) string {
	h := sha256.New()

	h.Write([]byte(password))

	return string(h.Sum(nil))
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
