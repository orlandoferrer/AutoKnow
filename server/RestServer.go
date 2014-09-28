package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlandoferrer/AutoKnow/controller"
	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

type RestServer struct {
	Port           int
	linkController controller.LinkController
}

func (restServer RestServer) Init() {
	if restServer.Port == 0 {
		restServer.Port = 8080
	}

	newLinkDao := &db.LinkDaoMap{}
	restServer.linkController = controller.NewLinkController(newLinkDao)
	restServer.linkController.Init()

	r := mux.NewRouter()
	r.HandleFunc("/edit/", editHandler(restServer))
	r.HandleFunc("/hello/", helloHandler(restServer))
	r.HandleFunc("/{linkResourcePath}", redirectHandler(restServer))

	r.HandleFunc("/api/newlink", newLinkHandler(restServer))
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%v", restServer.Port), nil)
}

type LinkService struct {
}

type LinkInfo struct {
	Message string `json:"message"`
}

func helloHandler(restServer RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(response, "hello!")
	}
}

func newLinkHandler(restServer RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		fmt.Printf("hello!\n")
		decoder := json.NewDecoder(req.Body)
		// var linkInfo LinkInfo
		var linkInfo model.Link
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
		linkInfo.InitLink()
		log.Printf("Created link:%v\n", linkInfo)
		restServer.linkController.CreateLink(linkInfo)
	}
}

func redirectHandler(restServer RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		linkFound := restServer.linkController.GetLinkByResourcePath(vars["linkResourcePath"])

		log.Printf("Redirecting to link:%v\n", linkFound)
		http.Redirect(response, req, linkFound.RedirectionPath, 302)
	}
}

func editHandler(restServer RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
	}
}
