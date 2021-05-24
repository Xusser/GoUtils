package goutils_test

import (
	"goutils"
	"testing"
)

func TestIsMAC(t *testing.T) {
	var macTests = []struct {
		in       string
		expected bool
	}{
		{"00:00:5e:00:53:01", true},
		{"02:00:5e:10:00:00:00:01", true},
		{"00:00:00:00:fe:80:00:00:00:00:00:00:02:00:5e:10:00:00:00:01", true},
		{"00-00-5e-00-53-01", true},
		{"02-00-5e-10-00-00-00-01", true},
		{"00-00-00-00-fe-80-00-00-00-00-00-00-02-00-5e-10-00-00-00-01", true},
		{"0000.5e00.5301", true},
		{"0200.5e10.0000.0001", true},
		{"0000.0000.fe80.0000.0000.0000.0200.5e10.0000.0001", true},
		{"01.02.03.04.05.06", false},
		{"01:02:03:04:05:06:", false},
		{"01002:03:04:05:06", false},
		{"01:02003:04:05:06", false},
		{"01:02:03004:05:06", false},
		{"01:02:03:04005:06", false},
		{"01:02:03:04:05006", false},
		{"01-02:03:04:05:06", false},
		{"01:02-03-04-05-06", false},
		{"0123:4567:89AF", false},
		{"0123-4567-89AF", false},
	}

	for _, tt := range macTests {
		actual := goutils.IsMAC(tt.in)
		if actual != tt.expected {
			t.Fatalf("Unexpected result:%v, expected:%v", actual, tt.expected)
		}
	}
}

func TestIsIP(t *testing.T) {
	var ipTests = []struct {
		in       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"19216811", false},
		{"114.514.19.19", false},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"::1", true},
	}

	for _, tt := range ipTests {
		_, actual := goutils.IsIP(tt.in)
		if actual != tt.expected {
			t.Fatalf("Unexpected result:%v, expected:%v", actual, tt.expected)
		}
	}
}
