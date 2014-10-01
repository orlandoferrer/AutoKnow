package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlandoferrer/AutoKnow/controller"
	"github.com/orlandoferrer/AutoKnow/db"
	"github.com/orlandoferrer/AutoKnow/model"
)

type RestServer struct {
	Port           int
	linkController *controller.LinkController
}

func (restServer *RestServer) Init() {
	if restServer.Port == 0 {
		restServer.Port = 8080
	}

	newLinkDao := &db.LinkDaoMap{}
	newController := controller.NewLinkController(newLinkDao)
	restServer.linkController = newController
	restServer.linkController.Init()

	r := mux.NewRouter()
	r.HandleFunc("/edit/", editHandler(restServer))
	r.HandleFunc("/hello/", helloHandler(restServer))
	r.HandleFunc("/{linkResourcePath}", redirectHandler(restServer))

	r.HandleFunc("/api/newlink", newLinkHandler(restServer))
	r.HandleFunc("/api/getlinks", getLinksHandler(restServer))
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%v", restServer.Port), nil)
}

type LinkService struct {
}

type LinkInfo struct {
	Message string `json:"message"`
}

func corsEnable(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	//fmt.Println("corsEnable")
}

func helloHandler(restServer *RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		corsEnable(&response)
		fmt.Fprintf(response, "hello!")
	}
}

func newLinkHandler(restServer *RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		corsEnable(&response)
		log.Printf("Req method:%v\n", req.Method)
		log.Printf("Request:%v\n", req)
		if "OPTIONS" == req.Method {

			return
		}
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

func getLinksHandler(restServer *RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		corsEnable(&response)

		links := restServer.linkController.GetResourcePathToLinks()
		responseMap := make(map[string][]model.Link)
		responseMap["links"] = links
		b, err := json.Marshal(responseMap)

		if err != nil {
			log.Fatalf("Errored out with:%v\n", err)
		}

		fmt.Fprintf(response, string(b))

	}
}

func redirectHandler(restServer *RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		corsEnable(&response)
		vars := mux.Vars(req)
		// log.Printf("Vars:%v, req:%v\n", vars, req)
		// log.Printf("Header:%v\n", req.Header)
		linkResourcePathToFind := vars["linkResourcePath"]
		linkFound, err := restServer.linkController.GetLinkByResourcePath(linkResourcePathToFind)
		if err != nil {
			log.Printf("Error trying to redirect:%v\n", err)
			http.Error(response, "Could not find extension with that link", 404)
			return
		}

		log.Printf("Using path %v found linke %v\n", linkResourcePathToFind, linkFound)
		// req.RemoteAddr
		log.Printf("Req:%v\n", req)
		// ipStringAll := strings.Split(req.RemoteAddr, ":")
		//ipString := ipStringAll[0]

		ipStringAll := req.RemoteAddr
		ipString, _, err := net.SplitHostPort(ipStringAll)
		if err != nil {
			log.Fatalf("Failing out splitting ip with:%v\n", err)
		}
		log.Printf("Req.RemoteAddr:%v\n", req.RemoteAddr)
		log.Printf("Req.RemoteAddrFirst:%v\n", ipString)
		newPageView := model.NewPageView(ipString, req.Header)
		linkFound.PageViews = append(linkFound.PageViews, *newPageView)

		restServer.linkController.UpdateLinkFoundByResourcePath(linkFound.ResourcePath,
			*linkFound)

		log.Printf("Redirecting to link:%v\n", linkFound)
		log.Printf("Total pageviews:%v\n", len(linkFound.PageViews))
		http.Redirect(response, req, linkFound.RedirectionPath, 302)
		// http.Redirect(response, req, "http://www.google.com", 302)
	}
}

func editHandler(restServer *RestServer) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		corsEnable(&response)
	}
}
