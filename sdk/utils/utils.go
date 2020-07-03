package utils

import (
	"crypto/hmac"
	"crypto/sha1"
)

func ShaHmac1(source, secret string) []byte {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(source))
	signedBytes := h.Sum(nil)
	return signedBytes
}

func StrArrIndex(xs []string, y string) int {
	for idx, x := range xs {
		if x == y {
			return idx
		}
	}
	return -1
}

func StrArrContains(xs []string, y string) bool {
	return StrArrIndex(xs, y) > -1
}

func StrArrContainsAll(xs, ys []string) bool {
	for _, x := range xs {
		if !StrArrContains(ys, x) {
			return false
		}
	}
	return true
}