package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		log.Printf("Header field %q, Value %q\n", k, v)
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api", homeLink)
	router.HandleFunc("/api/v1/applications", homeLink)
	router.HandleFunc("/v1/applications", homeLink)
	log.Fatal(http.ListenAndServe(":8100", router))
}
