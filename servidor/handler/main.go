package main

import (
	"log"
	"net/http"
)

func main() {
	redirect := http.RedirectHandler("https://www.google.com.pe", http.StatusMovedPermanently)
	notFound := http.NotFoundHandler()

	mux := http.NewServeMux()

	mux.Handle("/redirect", redirect)
	mux.Handle("/not", notFound)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
