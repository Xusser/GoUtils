package goutils

import (
	"net"
)

// IsIP 判断输入字符串是否为IP地址
func IsIP(addr string) (ip net.IP, ok bool) {
	ip = net.ParseIP(addr)
	ok = ip != nil
	return
}

// IsMAC 判断输入字符串是否为MAC地址
func IsMAC(mac string) (ok bool) {
	_, err := net.ParseMAC(mac)
	return err == nil
}
