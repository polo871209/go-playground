package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example:
// Authorization: ApiKey {ApiKey}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no key provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid format")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid format")
	}
	return vals[1], nil
}
