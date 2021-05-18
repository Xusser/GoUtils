package goutils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

// GetJSON 将Interface编码为JSON格式字符串 不会抛错
func GetJSON(value interface{}) string {
	result, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(result)
}

// GetStack 获取堆栈信息
func GetStack() string {
	buf := make([]byte, 4096)
	for {
		size := runtime.Stack(buf, true)
		if size == len(buf) {
			buf = make([]byte, len(buf)<<1) // buf*2
			continue
		}
		return string(buf[:size])
	}
}

// GetExecutablePath 获取应用程序目录
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// GetWorkPath 获取工作目录
func GetWorkPath() (string, error) {
	return os.Getwd()
}
