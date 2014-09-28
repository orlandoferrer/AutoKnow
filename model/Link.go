package model

import (
	"time"
)

//TODO: Priority Level: high priority level send email + SMS

type Link struct {
	ResourcePath     string          `json:"resourcePath"`
	PageViews        []PageView      `json:"pageViews"`
	CreationTime     time.Time       `json:"creationTime"`
	ExpirationTime   time.Time       `json:"expirationTime"`
	RedirectionPath  string          `json:"redirectionPath"`
	EnabledFlag      bool            `json:"enabledFlag"`
	IPRangeToLinkMap map[string]Link `json:"ipRangeToLinkMap"`
	// See here to see how to unmarshall this: http://stackoverflow.com/questions/21609363/unmarshalling-json-string-into-a-constant-in-google-go
	LinkType Type `json:"linkType"`
}

func (link *Link) InitLink() {
	if link.PageViews == nil {
		link.PageViews = make([]PageView, 0)
	}
	link.CreationTime = time.Now()
	if link.ExpirationTime.Before(link.CreationTime) {
		link.ExpirationTime = link.CreationTime.AddDate(1, 0, 0)
	}
	link.EnabledFlag = true
	if link.IPRangeToLinkMap == nil {
		link.IPRangeToLinkMap = make(map[string]Link)
	}
	if link.LinkType == DEFAULT {
		link.LinkType = GetResource
	}
}

type Type int

const (
	DEFAULT Type = iota
	Pixel
	GetResource
	HTML
	JS
)
