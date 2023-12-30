package otp

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	"strconv"
	"encoding/binary"
)

func Hotp(key []byte, counter uint64, digits int) int {
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func HotpStr(key string, counter int, digits int) (string, error) {
	n, err := strconv.ParseUint(string([]byte(key)[counter: counter + 20]), 10, 64)
	if err != nil {
		return "", err
	}
	n++
	raw, err := DecodeKey(key)
	if err != nil {
		return "", err
	}
	code := Hotp(raw, n, digits)
	codeStr := fmt.Sprintf("%0*d", digits, code)
	return codeStr, nil
}
