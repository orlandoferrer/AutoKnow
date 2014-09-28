package main

import (
	"log"

	"github.com/orlandoferrer/AutoKnow/server"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	restServer := &server.RestServer{}
	restServer.Init()
}
