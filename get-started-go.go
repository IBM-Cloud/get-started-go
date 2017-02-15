package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
)

var index = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/index.html",
))

func helloworld(w http.ResponseWriter, req *http.Request) {
  index.Execute(w, nil)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", helloworld)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
