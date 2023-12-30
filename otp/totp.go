package otp

import (
	"fmt"
	"time"
)

func Totp(key []byte, t time.Time, digits int) int {
	return Hotp(key, uint64(t.UnixNano())/30e9, digits)
}

func TotpStr(key string, digits int) (string, error) {
	raw, err := DecodeKey(key)
	if err != nil {
		return "", err
	}
	code := Totp(raw, time.Now(), digits)
	codeStr := fmt.Sprintf("%0*d", digits, code)
	return codeStr, nil
}
