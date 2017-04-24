package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := "localhost:4000"

	mux := http.NewServeMux()
	muxLogged := http.NewServeMux()

	muxLogged.HandleFunc("/v1/hello1", HelloHandler1)
	muxLogged.HandleFunc("/v1/hello2", logReqs(HelloHandler2))

	mux.HandleFunc("/v1/hello3", HelloHandler3)
	mux.Handle("/v1/", logRequests(muxLogged))

	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
