package db

import (
	"log"

	"github.com/orlandoferrer/AutoKnow/model"
)

type LinkDaoMap struct {
	linkMap map[string]model.Link
}

func (linkDao *LinkDaoMap) Init() error {
	linkDao.linkMap = make(map[string]model.Link)
	return nil
}

func (linkDao *LinkDaoMap) FindLinkByResourcePath(resourcePath string) *model.Link {
	linkFound := linkDao.linkMap[resourcePath]
	log.Printf("ResourcePath input:%v.  Found link:%v.  Whole map:%v.\n",
		resourcePath, linkFound, linkDao.linkMap)
	return &linkFound
}
func (linkDao *LinkDaoMap) CreateLink(link model.Link) error {
	linkDao.linkMap[link.ResourcePath] = link
	log.Printf("Created link %v\n", linkDao.linkMap[link.ResourcePath])
	return nil
}
func (linkDao *LinkDaoMap) UpdateLinkFoundByResourcePath(resourcePath string, link model.Link) error {
	linkDao.linkMap[resourcePath] = link
	return nil
}
