package db

import (
	"log"
	"testing"

	"github.com/orlandoferrer/AutoKnow/model"
)

func TestTestMapCreation(t *testing.T) {
	linkDaoMap := &LinkDaoMap{}
	linkDaoMap.Init()
	log.Printf("Before adding to it, map:%v\n", linkDaoMap)
	log.Printf("")
	linkDaoMap.CreateLink(model.Link{ResourcePath: "SomeName", RedirectionPath: "www.google.com"})
	findLink := linkDaoMap.FindLinkByResourcePath("SomeName")
	if findLink == nil {
		t.Error("Not inserting into map correctly.")
	}
	log.Printf("After adding to it, map:%v\n", linkDaoMap)
}
