package main

import (
	"net/http"
	"log"
	"runtime/debug"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	const templateName string = "index.html"

	context := Context{
		Firstname: "Boris",
	}

	jsonContext, err := NewJsonContext(context)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	template := Templates.Lookup(templateName)
	template.Execute(w, jsonContext)
}
