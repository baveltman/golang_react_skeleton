package main

import (
	"log"
	"net/http"
	"time"
	"os"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lbScheme := r.Header.Get(HEADER_LOAD_BALANCER_SCHEME)
		if lbScheme == "http" {
			log.Printf("redirecting %s to https", r.RequestURI)
			domain := os.Getenv(ENV_CANNONICAL_DOMAIN)
			redirect := domain + r.RequestURI
			http.Redirect(w, r, redirect, 301)
			return
		}

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
