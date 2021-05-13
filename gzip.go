package goutils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

// GzipCompressString GZIP压缩
func GzipCompressString(data []byte) (string, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data)
	if err != nil {
		return "", nil
	}

	if err := zw.Close(); err != nil {
		return "", nil
	}

	compressed, err := ioutil.ReadAll(&buf)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(compressed), nil
}

// GzipDecompressString GZIP解压
func GzipDecompressString(data string) (string, error) {
	result := ""
	z, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return result, err
	}
	r, err := gzip.NewReader(bytes.NewReader(z))
	if err != nil {
		return result, err
	}
	byteResult, err := ioutil.ReadAll(r)
	if err != nil {
		return result, err
	}
	result = string(byteResult)
	return result, nil
}
