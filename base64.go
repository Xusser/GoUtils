package goutils

import (
	"encoding/base64"
	"fmt"
)

// DecodeBase64String 解码经Base64编码的String
func DecodeBase64String(str string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode base64 string! %s", err)
	}
	return decodedBytes[:], nil
}

// EncodeBase64 对Bytes进行Base64编码
func EncodeBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// EncodeBase64String 对Bytes进行Base64编码
func EncodeBase64String(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// EncodeURLSafeBase64 将Bytes使用Base64编码(URL安全)
func EncodeURLSafeBase64(bytes []byte) string {
	return base64.RawURLEncoding.EncodeToString(bytes)
}

// DecodeURLSafeBase64 对String使用Base64解码(URL安全)
func DecodeURLSafeBase64(str string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(str)
}
