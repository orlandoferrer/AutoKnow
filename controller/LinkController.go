package controller

import (
	"errors"

	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

type LinkController struct {
	linkDao db.LinkDao
}

func NewLinkController(linkDao db.LinkDao) LinkController {
	return LinkController{linkDao}
}

func (linkController LinkController) Init() error {
	//TODO Do actual error handling
	return linkController.linkDao.Init()
}

func (linkController LinkController) CreateLink(link model.Link) error {
	//TODO Do actual error handling
	exists := linkController.doesResourceExist(link)
	if exists == false {

		linkController.linkDao.CreateLink(link)
	} else {
		return errors.New("Link already exists.")
	}

	return nil
}

func (linkController LinkController) doesResourceExist(link model.Link) bool {
	//TODO Do actual error handling
	linkFound := linkController.linkDao.FindLinkByResourcePath(link.ResourcePath)
	return linkFound != nil
}

func (linkController LinkController) GetLinkByResourcePath(resourcePath string) *model.Link {
	return linkController.linkDao.FindLinkByResourcePath(resourcePath)
}
