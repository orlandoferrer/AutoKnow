package model

import (
	"net"
	"net/http"
	"time"
)

type PageView struct {
	CreationTime time.Time   `json:"creationTime"`
	IPAddress    net.IP      `json:"ipAddress"`
	HTTPHeaders  http.Header `json:"httpHeaders"`
}

func NewPageView(ipString string, headers http.Header) *PageView {
	return &PageView{
		CreationTime: time.Now(),
		IPAddress:    net.ParseIP(ipString),
		HTTPHeaders:  headers,
	}
}
