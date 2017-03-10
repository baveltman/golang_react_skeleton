package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

const DEFAULT_PORT string = "8080"

var Templates *template.Template

var StaticFileServer http.Handler

var ApiServerHost string

func main() {

	//initialize templates
	Templates = InitTemplates()

	//initialize static file server
	StaticFileServer = InitStaticAssetServer()

	//initialize general router
	router := NewRouter()

	//choose port
	port := os.Getenv(ENV_PORT)
	if port == "" {
		port = DEFAULT_PORT
	}

	//listen and serve
	log.Println("Web Client Server listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
