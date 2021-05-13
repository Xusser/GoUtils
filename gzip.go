package goutils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

func GzipCompress(data []byte) (compressed []byte, err error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err = zw.Write(data); err != nil {
		return
	}

	if err = zw.Close(); err != nil {
		return
	}

	compressed, err = ioutil.ReadAll(&buf)
	return
}

func GzipDecompress(compressed []byte) (data []byte, err error) {
	r, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GzipCompressString GZIP压缩
func GzipCompressString(data string) (string, error) {
	compressed, err := GzipCompress([]byte(data))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(compressed), nil
}

// GzipDecompressString GZIP解压
func GzipDecompressString(data string) (string, error) {
	compressed, err := DecodeBase64(data)
	if err != nil {
		return "", err
	}

	b, err := GzipDecompress(compressed)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
