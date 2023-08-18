package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func ParseHeaderAuthToken(key string, r *http.Request) (string, error) {

	tok := r.Header.Get("Authorization")
	tok = strings.TrimSpace(tok)

	if tok == "" {
		return "", fmt.Errorf("auth token is empty")
	}

	value := strings.Split(tok, " ")

	if len(value) < 2 || value[0] != key {
		return "", fmt.Errorf("invalid token format in header authorization")
	}

	if len(value[1]) == 0 {
		return "", fmt.Errorf("empty auth token")
	}

	return value[1], nil
}
