package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "General\n")
	fmt.Fprintf(w, "Request URL: %q\n", r.RequestURI)
	fmt.Fprintf(w, "Request Method: %q\n", r.Method)
	fmt.Fprintf(w, "Remote Address: %q\n", r.RemoteAddr)

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Request Headers\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err == nil && len(body) > 0 {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Body\n")
		fmt.Fprintf(w, "%s", string(body))
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").HandlerFunc(defaultHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("http-echo listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
