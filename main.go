package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "General\n")
	fmt.Fprintf(w, "Request URL: %q\n", r.RequestURI)
	fmt.Fprintf(w, "Request Method: %q\n", r.Method)
	fmt.Fprintf(w, "Request http protocol: %q\n", r.Proto)
	fmt.Fprintf(w, "Remote Address: %q\n", r.RemoteAddr)

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Request Headers\n")
	headers := make([]string, 0, len(r.Header))
	for k := range r.Header {
		headers = append(headers, k)
	}
	sort.Strings(headers)

	for _, k := range headers {
		fmt.Fprintf(w, "%q: %q\n", k, r.Header[k])
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

	// for graceful shutdown of service.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Created a new server instance
	log.Printf("[http-echo] listening on port %s", port)
	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("[http-echo]", "failed,", err)
		}
	}()

	<-done

	log.Println("[Grcefull shutdown]")
	// Gracefull Shutdown added.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

}
