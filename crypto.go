package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// GenerateAESNonce 生成AES使用的Nonce
func GenerateAESNonce() ([]byte, error) {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

// AES256GCMEncrypt AES-256-GCM 加密
func AES256GCMEncrypt(data []byte, key string, nonce []byte) ([]byte, error) {
	// 32bytes -> AES-256
	hash := sha256.Sum256([]byte(key))
	hashedKey := hash[:]

	block, err := aes.NewCipher([]byte(hashedKey))
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	result := aesgcm.Seal(nil, nonce, data, nil)

	return result, nil
}

// AES256GCMDecrypt AES-256-GCM 解密
func AES256GCMDecrypt(data []byte, key string, nonce []byte) ([]byte, error) {
	// 32bytes -> AES-256
	hash := sha256.Sum256([]byte(key))
	hashedKey := hash[:]

	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	result, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
