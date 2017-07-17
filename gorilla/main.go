package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nombre := vars["nombre"]
	id := vars["id"]
	w.Write([]byte("Gorilla!\n"))
	fmt.Fprintf(w, "\nLos parametros son "+nombre+", "+id)
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/usuarios/{id:[0-9]+}", YourHandler).Methods("PUT", "DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
