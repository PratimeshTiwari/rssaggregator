package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authoirzation : ApiKey {apikey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing API key in Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid header, expected 'ApiKey {apikey}'")
	}

	return vals[1], nil

}
