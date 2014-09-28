package db

import (
	"github.com/orlandoferrer/AutoKnow/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LinkDaoMongo struct {
}

type LinkDao interface {
	Init() error
	FindLinkByResourcePath(resourcePath string) *model.Link
}
