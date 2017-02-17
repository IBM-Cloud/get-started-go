package main

import (
	"log"
	"os"
	"net/http"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

import "github.com/timjacobi/go-couchdb"

type Visitor struct {
    Name      string    `json:"name"`
}

type Visitors []Visitor

type alldocsResult struct {
	TotalRows int `json:"total_rows"`
	Offset    int
	Rows      []map[string]interface{}
}

func main() {
	r := gin.Default()

	r.StaticFile("/", "./static/index.html")

	r.Static("/static", "./static")

	var dbName = "mydb"

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file does not exist")
	}

  cloudantUrl := os.Getenv("CLOUDANT_URL")

	appEnv, _ := cfenv.Current()
  _ = appEnv
  if(appEnv!=nil){
    cloudantService, _ := appEnv.Services.WithLabel("cloudantNoSQLDB")
    if(len(cloudantService)>0){
      cloudantUrl = cloudantService[0].Credentials["url"].(string)
    }
  }

  cloudant, err := couchdb.NewClient(cloudantUrl, nil)
	if err != nil {
		log.Println("error Cloudant NewClient")
	}

  //ensure db exists
  //if the db exists the db will be returned anyway
  cloudant.CreateDB(dbName)

	r.POST("/api/visitors", func(c *gin.Context) {
		var visitor Visitor

    if c.BindJSON(&visitor) == nil {
      cloudant.DB(dbName).Post(visitor)

			c.String(200, "Hello "+visitor.Name)
		}

	})

  r.GET("/api/visitors", func(c *gin.Context) {
    var result alldocsResult

    if cloudantUrl == "" {
      c.JSON(200, gin.H{})
      return
    }

    err := cloudant.DB(dbName).AllDocs(&result, couchdb.Options{"include_docs": true})
    if err != nil {
      log.Println("error Cloudant AllDocs")
      c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch docs"})
    } else {
      c.JSON(200, result.Rows)

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
