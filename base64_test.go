package goutils_test

import (
	"goutils"
	"testing"
)

func TestDecodeBase64String(t *testing.T) {
	result, err := goutils.DecodeBase64String("dGVzdA==")
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if string(result) != "test" {
		t.Fatal("Unexpected result:", result)
	}
	t.Log("OK")
}

func TestEncodeBase64(t *testing.T) {
	result := goutils.EncodeBase64([]byte("test"))

	if result != "dGVzdA==" {
		t.Fatal("Unexpected result:", result)
	}
	t.Log("OK")
}

func TestEncodeBase64String(t *testing.T) {
	result := goutils.EncodeBase64String("test")

	if result != "dGVzdA==" {
		t.Fatal("Unexpected result:", result)
	}
	t.Log("OK")
}

func TestEncodeURLSafeBase64(t *testing.T) {
	result := goutils.EncodeURLSafeBase64([]byte("test"))

	if result != "dGVzdA" {
		t.Fatal("Unexpected result:", result)
	}
	t.Log("OK")
}

func TestDecodeURLSafeBase64(t *testing.T) {
	result, err := goutils.DecodeURLSafeBase64("dGVzdA")
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if string(result) != "test" {
		t.Fatal("Unexpected result:", result)
	}
	t.Log("OK")
}
