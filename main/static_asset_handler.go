package main

import (
	"net/http"
	"os"
	"fmt"
)

func InitStaticAssetServer() http.Handler{
	var staticFilesPath string = ""
	isDocker := os.Getenv(ENV_DOCKER)
	if isDocker != "" {
		pwd := os.Getenv(ENV_DIR_ENTRY)
		if pwd == "" {
			fmt.Println("Missing ENV " + ENV_DIR_ENTRY)
			os.Exit(1)
		}
		staticFilesPath = pwd + "/" + DIR_STATIC_ASSETS_FOLDER + "/"
	} else {
		staticFilesPath = DIR_STATIC_ASSETS_FOLDER
	}

	return http.FileServer(http.Dir(staticFilesPath))

}

func StaticAssetHandler(w http.ResponseWriter, r *http.Request) {
	handler := http.StripPrefix("/static/", StaticFileServer)
	handler.ServeHTTP(w, r)
}
