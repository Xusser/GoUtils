package goutils

import (
	"encoding/base64"
	"fmt"
)

// EncodeBase64 Bytes编码为Base64
func EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// DecodeBase64 Base64解码为Bytes
func DecodeBase64(str string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, fmt.Errorf("unable to decode base64 string! %s", err)
	}
	return decodedBytes[:], nil
}

// EncodeBase64String String编码为Base64
func EncodeBase64String(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// DecodeBase64String Base64解码为String
func DecodeBase64String(src string) (string, error) {
	b, err := DecodeBase64(src)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// EncodeURLSafeBase64 Bytes编码为Base64(URL安全)
func EncodeURLSafeBase64(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}

// DecodeURLSafeBase64 Base64(URL安全)解码为Bytes
func DecodeURLSafeBase64(src string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(src)
}

// EncodeURLSafeBase64String String编码为Base64(URL安全)
func EncodeURLSafeBase64String(src string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(src))
}

// DecodeURLSafeBase64String Base64(URL安全)解码为String
func DecodeURLSafeBase64String(src string) (dst string, err error) {
	b, err := base64.RawURLEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
