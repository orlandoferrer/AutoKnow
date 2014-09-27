package model

import (
	"bytes"
	"net"
	"strings"
)

type IPRange struct {
	From net.IP
	To   net.IP
}

func NewIPRange(ipRange string) IPRange {
	strippedIPRange := strings.Replace(ipRange, " ", "", -1)
	slicedRange := strings.Split(strippedIPRange, "-")
	//TODO input validation

	return IPRange{
		From: net.ParseIP(slicedRange[0]),
		To:   net.ParseIP(slicedRange[1]),
	}
}

func (ipRange IPRange) IsInRange(ipString string) bool {
	parsedIP := net.ParseIP(ipString)
	if bytes.Compare(parsedIP, ipRange.From) >= 0 && bytes.Compare(parsedIP, ipRange.To) <= 0 {
		return true
	} else {
		return false
	}
}
