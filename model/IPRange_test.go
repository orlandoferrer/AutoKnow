package model

import (
	"net"
	"testing"
)

func TestNewIPRange(t *testing.T) {
	testIPStrings := []string{
		"172.32.43.56  -  172.32.44.31",
		"168.234.54.32 - 168.235.134.6",
		"193.254.77.43-193.254.77.200",
		"164.300.76.23 - 164.301.56.67",
	}

	for i := range testIPStrings {

	}

}

var rangePair struct {
	IPAddressRange IPRange
	ipAddress      string
}

func TestIsInRange(t *testing.T) {
	var tests = []rangePair{
		{IPRange{From: net.ParseIP("172.135.23.21"), To: net.ParseIP("172.137.56.24")}, net.ParseIP("172.136.56.87")},
		{IPRange{From: net.ParseIP("125.135.23.21"), To: net.ParseIP("129.137.56.24")}, net.ParseIP("124.136.56.87")},
		{IPRange{From: net.ParseIP("172.135.23.21"), To: net.ParseIP("172.137.56.24")}, net.ParseIP("172.136.56.87")},
		{IPRange{From: net.ParseIP("172.135.23.21"), To: net.ParseIP("172.137.56.24")}, net.ParseIP("172.136.56.87")},
		{IPRange{From: net.ParseIP("172.135.23.21"), To: net.ParseIP("172.137.56.24")}, net.ParseIP("172.136.56.87")},
	}

}
