package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Inedx Page",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"Static Assets JS",
		"GET",
		"/static/dist/js/{fileName}",
		StaticAssetHandler,
	},
	Route{
		"Static Assets CSS",
		"GET",
		"/static/dist/css/{fileName}",
		StaticAssetHandler,
	},
	Route{
		"Static Image Assets",
		"GET",
		"/static/img/{folder}/{fileName}",
		StaticAssetHandler,
	},
}
