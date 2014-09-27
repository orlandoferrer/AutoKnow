package model

import (
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
