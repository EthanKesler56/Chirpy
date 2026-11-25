package auth

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func MakeRefreshToken() string {
	buffer := make([]byte, 32)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Fatal("An error has occurred")

	}
	return hex.EncodeToString(buffer)

}
