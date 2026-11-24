package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authentication header")
	}
	authParts := strings.Fields(authHeader)
	if len(authParts) < 2 || authParts[0] != "Bearer" {
		return "", errors.New("malformed header")
	}
	return authParts[1], nil

}
