package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type zip struct {
	Zip   string `json:"zip"`
	City  string `json:"city"`
	State string `json:"state"`
}

/* a slice of pointers to zip */
type zipSlice []*zip

/* map of city to zip-codes */
type zipIndex map[string]zipSlice

/* The * character is used to define a pointer in both C and Go.
Instead of a real value the variable instead has an address to
the location of a value. */
func helloHandler(w http.ResponseWriter, r *http.Request) {
	/* Handling query strings.
	Get value of the name parameter. */
	name := r.URL.Query().Get("name")

	/* Declaring Content-Type. */
	w.Header().Add("Content-Type", "text/plain")

	/* Send as byte slice, so we can send pictures too.
	Slice = Java Arraylist.
	Automatically adds Content-Length for you because it knows exactly how many bytes. */
	w.Write([]byte("hello " + name))
}

/* We will be receiving a pointer, rather than a giant-ass struct.
ResponseWriter is an interface, and they are always passed by reference. */

/* this */
func (zi zipIndex) zipsForCityHandler(w http.ResponseWriter, r *http.Request) {
	// /zips/city/seattle is the format

	_, city := path.Split(r.URL.Path)
	lcity := strings.ToLower(city)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	/* Writes to the ResponseWriter */
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(zi[lcity]); err != nil { /* slice of zip struct */
		http.Error(w, "error encoding json: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	/* Declare and assign.
	All variables are statically typed in Go.
	:= create a new variable named addr and figure out what the type is by looking at
	whatever is on the right. */
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		/* log.Fatal writes the message to the console with the exit status 1. */
		log.Fatal("please set ADDR environment variable")
	}

	f, err := os.Open("../data/zips.json")

	if err != nil {
		log.Fatal("error opening zips file: " + err.Error())
	}

	/* A slice of pointers. make() only allocates. */
	/* length, capacity */
	zips := make(zipSlice, 0, 43000)
	decoder := json.NewDecoder(f)

	if err := decoder.Decode(&zips); err != nil {
		log.Fatal("error decoding zips json: " + err.Error())
	}

	fmt.Printf("loaded %d zips \n", len(zips))

	zi := make(zipIndex)

	/* range returns the index and the value at that index.
	We are using the underscore here because we don't need it.
	For each loop in Java */
	for _, z := range zips {
		lower := strings.ToLower(z.City)

		/* find zipcodes in zips that belong to Seattle and map it. */
		zi[lower] = append(zi[lower], z)
	}

	fmt.Printf("there are %d zips in Seattle\n", len(zi["seattle"]))

	/* When https calls are made on the /hello path, call helloHandler.
	Passing a pointer to the function.
	Registering helloHandler to this resource path. */
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/zips/city/", zi.zipsForCityHandler) /* set the Receiver to zi. */

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
