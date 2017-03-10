package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func InitTemplates() *template.Template {
	var allFiles []string

	pwd := os.Getenv(ENV_DIR_ENTRY)
	if pwd == "" {
		fmt.Println("Missing ENV " + ENV_DIR_ENTRY)
		os.Exit(1)
	}
	templatesDir := pwd + DIR_TEMPLATES

	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	templates := template.Must(template.ParseFiles(allFiles...))

	return templates
}
