package goutils_test

import (
	"goutils"
	"testing"
)

func TestAES256GCMEncrypt(t *testing.T) {
	nonce := "123456789012"
	key := "114514"
	src := "test"

	result, err := goutils.AES256GCMEncrypt([]byte(src), key, []byte(nonce))
	if err != nil {
		t.Fatal("Unexpected fail to encrypt, err:", err)
	}

	b := goutils.EncodeBase64(result)
	if b != "CMEhQLfT0rZBBFE6BkNQbbQR96g=" {
		t.Fatal("Unexpected base64 encoded result:", b)
	}
	t.Log("OK")
}
