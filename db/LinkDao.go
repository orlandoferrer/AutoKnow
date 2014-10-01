package db

import "github.com/orlandoferrer/AutoKnow/model"

// LinkDao is the generic interface for accessing any database for the Link object.
type LinkDao interface {
	Init() error
	// FindLinkByResourcePath returns Link from the DB based on resource path given.
	// If not found, returns null
	FindLinkByResourcePath(resourcePath string) (*model.Link, error)
	CreateLink(link model.Link) error
	UpdateLinkFoundByResourcePath(resourcePath string, link model.Link) error
	GetResourcePathToLinkMap() []model.Link
}
