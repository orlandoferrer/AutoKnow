package model

import (
	"time"
)

//TODO: Priority Level: high priority level send email + SMS

type Link struct {
	ResourcePath     string
	PageViews        []PageView
	CreationTime     time.Time
	ExpirationTime   time.Time
	RedirectionPath  string
	EnabledFlag      bool
	IPRangeToLinkMap map[string]Link
	LinkType         Type
}

type Type int

const (
	Pixel Type = iota
	GetResource
	HTML
	JS
)
