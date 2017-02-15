package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Visitor struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	r.StaticFile("/", "./static/index.html")

	r.Static("/static", "./static")

	r.POST("/api/visitors", func(c *gin.Context) {
		var json Visitor
		if c.BindJSON(&json) == nil {
			c.String(200, "Hello "+json.Name)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

// package main
//
// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"html/template"
// 	//for extracting service credentials from VCAP_SERVICES
// 	//"github.com/cloudfoundry-community/go-cfenv"
// )
//
// const (
// 	DEFAULT_PORT = "8080"
// )
//
// var index = template.Must(template.ParseFiles(
//   "templates/_base.html",
//   "templates/index.html",
// ))
//
// func helloworld(w http.ResponseWriter, req *http.Request) {
//   index.Execute(w, nil)
// }
//
// func main() {
// 	var port string
// 	if port = os.Getenv("PORT"); len(port) == 0 {
// 		port = DEFAULT_PORT
// 	}
//
// 	http.HandleFunc("/", helloworld)
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
//
// 	log.Printf("Starting app on port %+v\n", port)
// 	http.ListenAndServe(":"+port, nil)
// }
