package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RestServer struct {
	Port int
}

func (restServer RestServer) Init() {
	if restServer.Port == 0 {
		restServer.Port = 8080
	}
	r := mux.NewRouter()
	r.HandleFunc("/edit/", editHandler)
	r.HandleFunc("/hello/", helloHandler)
	r.HandleFunc("/{linkResourcePath}", linkHandler)

	r.HandleFunc("/api/newlink", newLinkHandler)
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%v", restServer.Port), nil)
}

type LinkService struct {
}

type LinkInfo struct {
	Message string `json:"message"`
}

func helloHandler(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "hello!")
}

func newLinkHandler(response http.ResponseWriter, req *http.Request) {
	fmt.Printf("hello!\n")
	decoder := json.NewDecoder(req.Body)
	var linkInfo LinkInfo
	err := decoder.Decode(&linkInfo)
	if err != nil {
		log.Fatalf("Errored out with:%v\n", err)
	}
	fmt.Printf("Message recieved: %v\n", linkInfo)
	fmt.Fprintf(response, "Replying!")
	//Parse input to get link information.
	// req
	//Use link information to create new link.
	//Return status information.
}

func linkHandler(response http.ResponseWriter, req *http.Request) {

}

func editHandler(response http.ResponseWriter, req *http.Request) {

}
