package controller

import (
	"log"
	"testing"

	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestController(t *testing.T) {

	newLinkDao := &db.LinkDaoMap{}

	newController := NewLinkController(newLinkDao)
	controllerPointer := newController
	controllerPointer.Init()

	newLink := model.Link{ResourcePath: "SomeName", RedirectionPath: "www.google.com"}
	log.Printf("Link about to create:%v\n", newLink)
	err := controllerPointer.CreateLink(newLink)
	if err != nil {
		t.Errorf("Error creating link:%v\n", err)
	}
	foundLink, err := controllerPointer.GetLinkByResourcePath(newLink.ResourcePath)
	if err != nil {
		t.Errorf("Error getting link:%v\n", err)
	}
	log.Printf("Link found:%v\n", foundLink)
}
