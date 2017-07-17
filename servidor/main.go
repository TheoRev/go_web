package main

import (
	"fmt"
	"log"
	"net/http"
)

// Hola Handler
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo.")
}

// HolaDos Handler
func HolaDos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo dos.")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/dos", HolaDos)

	// DefaultServeMux
	// http.HandleFunc("/", Hola)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
