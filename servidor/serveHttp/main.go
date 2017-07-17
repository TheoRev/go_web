package main

import (
	"fmt"
	"log"
	"net/http"
)

// User struct de usuarios
type User struct {
	name string
}

// ServeHTTP servicio http
func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola: "+this.name)
}

func main() {
	theo := &User{name: "Theo"}

	mux := http.NewServeMux()
	mux.Handle("/theo", theo)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
