package main

import "github.com/orlandoferrer/AutoKnow/server"

func main() {
	restServer := &server.RestServer{}
	restServer.Init()
}
