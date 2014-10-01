package controller

import (
	"errors"
	"log"

	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

type LinkController struct {
	linkDao *db.LinkDaoMap
}

func NewLinkController(linkDao *db.LinkDaoMap) *LinkController {
	return &LinkController{linkDao}
}

func (linkController *LinkController) Init() error {
	//TODO Do actual error handling
	return linkController.linkDao.Init()
}

func (linkController *LinkController) CreateLink(link model.Link) error {
	//TODO Do actual error handling
	exists := linkController.doesResourceExist(link)
	log.Printf("Exists?:%v", exists)
	if !exists {
		log.Printf("About to creat link from controllor:%v\n", link)
		err := linkController.linkDao.CreateLink(link)
		if err != nil {
			log.Printf("Not able to create link:%v", err)
		}

	} else {
		log.Printf("Issue trying to create link.")
		return errors.New("Link already exists.")
	}

	return nil
}

func (linkController *LinkController) doesResourceExist(link model.Link) bool {
	//TODO Do actual error handling
	linkFound, error := linkController.linkDao.FindLinkByResourcePath(link.ResourcePath)
	if error == nil {
		log.Printf("Found link when checking resource exists:%v", linkFound)
	}
	return error == nil
}

func (linkController *LinkController) UpdateLinkFoundByResourcePath(resourcePath string, link model.Link) error {
	return linkController.linkDao.UpdateLinkFoundByResourcePath(resourcePath, link)
}

func (linkController *LinkController) GetLinkByResourcePath(resourcePath string) (*model.Link, error) {
	return linkController.linkDao.FindLinkByResourcePath(resourcePath)
}

func (linkController *LinkController) GetResourcePathToLinks() []model.Link {
	return linkController.linkDao.GetResourcePathToLinks()
}
