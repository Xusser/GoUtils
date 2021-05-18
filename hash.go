package goutils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/crc64"
	"strings"
)

// CRC64 计算CRC64哈希值
func CRC64(payload []byte) string {
	h := crc64.New(crc64.MakeTable(crc64.ISO))
	h.Write(payload)
	return fmt.Sprintf("%x", h.Sum64())
}

// SHA256 计算SHA256哈希值
func SHA256(payload []byte) string {
	hash := sha256.Sum256(payload)
	result := hex.EncodeToString(hash[:])
	return strings.ToUpper(result)
}

// HMACSHA256 计算HMAC SHA256值
func HMACSHA256(payload, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(payload)
	return h.Sum(nil)
}

// HMACSHA1 计算HMAC SHA1值
func HMACSHA1(payload, secret []byte) []byte {
	h := hmac.New(sha1.New, secret)
	h.Write(payload)
	return h.Sum(nil)
}
