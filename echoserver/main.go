// Package math serves a simple web server with routing to different echo endpoints
package main

import (
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func ping(w http.ResponseWriter, r *http.Request) {
	message := "Ping"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func errorz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Reserved for future errors"))
}

func main() {
	h := http.NewServeMux()
	h.HandleFunc("/errorz", errorz)
	h.HandleFunc("/", hello)
	h.HandleFunc("/ping", ping)
	// http.Handle("/", r)

	err := http.ListenAndServe(":80", h)
	panic(err)
}
