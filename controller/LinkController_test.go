package controller

import (
	"log"
	"testing"

	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

func TestController(t *testing.T) {

	newLinkDao := &db.LinkDaoMap{}

	newController := NewLinkController(newLinkDao)
	controllerPointer := newController
	controllerPointer.Init()

	newLink := model.Link{ResourcePath: "SomeName", RedirectionPath: "www.google.com"}
	log.Printf("Link about to create:%v\n", newLink)
	controllerPointer.CreateLink(newLink)
	foundLink := controllerPointer.GetLinkByResourcePath(newLink.ResourcePath)
	log.Printf("Link found:%v\n", foundLink)
}
