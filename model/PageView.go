package model

import (
	"net"
	"time"
)

type PageView struct {
	CreationTime time.Time
	IPAddress    net.IP
	HTTPHeaders  map[string]string
}
