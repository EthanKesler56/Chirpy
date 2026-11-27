package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {

	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		err := errors.New("authorization header not found")
		return "", err
	}
	if !(strings.HasPrefix(authHeader, "ApiKey ")) {
		err := errors.New("Malformed header")
		return "", err
	}
	trimmedAuthHeader := strings.TrimPrefix(authHeader, "ApiKey ")
	return trimmedAuthHeader, nil
}
