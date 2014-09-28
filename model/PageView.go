package model

import (
	"net"
	"time"
)

type PageView struct {
	CreationTime time.Time         `json:"creationTime"`
	IPAddress    net.IP            `json:"ipAddress"`
	HTTPHeaders  map[string]string `json:"httpHeaders"`
}
