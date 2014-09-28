package controller

import (
	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

type LinkController struct {
	linkDao db.LinkDao
}

func (linkController LinkController) Init() error {
	//TODO Do actual error handling
	return linkController.linkDao.Init()
}

func (linkController LinkController) CreateLink(link model.Link) error {
	//TODO Do actual error handling
	return nil
}

func (linkController LinkController) isResourceAlreadyInUse(link model.Link) bool {
	//TODO Do actual error handling
	linkFound := linkController.linkDao.FindLinkByResourcePath(link.ResourcePath)
	return linkFound != nil
}
