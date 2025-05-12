package butils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHex(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic("unable to generate random bytes")
	}
	return hex.EncodeToString(b)
}
