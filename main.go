package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method %q, Request %q\n", r.Method, r.RequestURI)
	fmt.Printf("Method %q, Request %q\n", r.Method, r.RequestURI)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").HandlerFunc(homeLink)
	log.Printf("http-echo listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
