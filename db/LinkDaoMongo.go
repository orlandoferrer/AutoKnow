package db

import "github.com/orlandoferrer/AutoKnow/model"

type LinkDaoMongo struct {
}

type LinkDao interface {
	Init() error
	FindLinkByResourcePath(resourcePath string) *model.Link
}
