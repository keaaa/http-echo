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
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	addRequestedInformation(w, r)

	includeHostInfo := os.Getenv("INCLUDE_HOST_INFORMATION")
	if includeHostInfo == "true" || includeHostInfo == "1" {
		addHostingInformation(w, r)
	}
}

func addHostingInformation(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Hostname: %q\n", hostName)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Environment variable\n")
	envVar := os.Environ()
	for _, v := range envVar {
		keyVal := strings.Split(v, "=")
		fmt.Fprintf(w, "%q: %q\n", keyVal[0], keyVal[1])
	}
}

func addRequestedInformation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "General\n")
	fmt.Fprintf(w, "Request URL: %q\n", r.RequestURI)
	fmt.Fprintf(w, "Request Method: %q\n", r.Method)
	fmt.Fprintf(w, "Request HTTP protocol: %q\n", r.Proto)
	fmt.Fprintf(w, "Remote Address: %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Request Host: %q\n", r.Host)
	fmt.Fprintf(w, "Request content length: %v\n", r.ContentLength)
	fmt.Fprintf(w, "Request TransferEncoding: %q\n", r.TransferEncoding)

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
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Body\n")
	if err == nil && len(body) > 0 {
		fmt.Fprintf(w, "%s", string(body))
	}

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Response headers\n")
	responseHeaders := w.Header()
	headers = make([]string, 0, len(responseHeaders))
	for k := range responseHeaders {
		headers = append(headers, k)
	}
	sort.Strings(headers)

	for _, k := range headers {
		fmt.Fprintf(w, "%q: %q\n", k, responseHeaders[k])
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
