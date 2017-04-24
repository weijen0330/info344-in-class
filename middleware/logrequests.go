package main

import (
	"log"
	"net/http"
	"time"
)

func logReq(r *http.Request) {
	log.Println(r.Method, r.URL.Path)
}

// Closure: the inner function can use the parameter from the outer function
func logReqs(hfn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		start := time.Now()
		hfn(w, r)
		// %v will return human readable format
		log.Printf("%v\n", time.Since(start))
	}
}

// Taking in a full Handler interface and return a full Handler interface on the other
// side as well.
func logRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s", r.Method, r.URL.Path)
		start := time.Now()
		handler.ServeHTTP(w, r)
		logger.Printf("%v\n", time.Since(start))
	})
}
