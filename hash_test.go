package goutils_test

import (
	"encoding/hex"
	"goutils"
	"testing"
)

func TestCRC64(t *testing.T) {
	source := []byte("1145141919810")
	expect := "5b5ef7876b88e364"
	result := goutils.CRC64(source)
	if result != expect {
		t.Fatal("Unexpected result:", result, ", expected:", expect)
	}
}

func BenchmarkCRC64(b *testing.B) {
	source := []byte("1145141919810")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goutils.CRC64(source)
	}
}

func TestSHA256(t *testing.T) {
	source := []byte("1145141919810")
	expect := "b7ab30a912521ac36e433a5cfc8b5c1037884487af45ae5311ced235ee77faef"
	result := goutils.SHA256(source)
	if result != expect {
		t.Fatal("Unexpected result:", result, ", expected:", expect)
	}
}

func BenchmarkSHA256(b *testing.B) {
	source := []byte("1145141919810")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goutils.SHA256(source)
	}
}

func TestHMACSHA256(t *testing.T) {
	source := []byte("1145141919810")
	secret := []byte("AreUOK")
	expect := "1aff4c104640bddab2cbee14adea94984c283ff6f77a66f05e7a95226ee75026"
	result := hex.EncodeToString(goutils.HMACSHA256(source, secret))
	if result != expect {
		t.Fatal("Unexpected result:", result, ", expected:", expect)
	}
}

func BenchmarkHMACSHA256(b *testing.B) {
	source := []byte("1145141919810")
	secret := []byte("AreUOK")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goutils.HMACSHA256(source, secret)
	}
}

func TestHMACSHA1(t *testing.T) {
	source := []byte("1145141919810")
	secret := []byte("AreUOK")
	expect := "a3965308303b646c0472bf11ac300f7fdfa44fa2"
	result := hex.EncodeToString(goutils.HMACSHA1(source, secret))
	if result != expect {
		t.Fatal("Unexpected result:", result, ", expected:", expect)
	}
}

func BenchmarkHMACSHA1(b *testing.B) {
	source := []byte("1145141919810")
	secret := []byte("AreUOK")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goutils.HMACSHA1(source, secret)
	}
}
