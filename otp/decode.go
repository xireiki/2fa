package otp

import (
	"encoding/base32"
	"strings"
)

func DecodeKey(key string) ([]byte, error) {
	raw, err := base32.StdEncoding.DecodeString(strings.ToUpper(key))
	if err != nil {
		return nil, err
	}
	return raw, nil
}