package otp

import (
	"fmt"
	"time"
)

func Totp(key []byte, t time.Time, digits int, period int) int {
	return Hotp(key, uint64(t.UnixNano())/(uint64(period) * 1e9), digits)
}

func TotpStr(key string, digits int, period int) (string, error) {
	raw, err := DecodeKey(key)
	if err != nil {
		return "", err
	}
	code := Totp(raw, time.Now(), digits, period)
	codeStr := fmt.Sprintf("%0*d", digits, code)
	return codeStr, nil
}
