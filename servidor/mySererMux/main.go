package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheoRev/go_web/servidor/mySererMux/projects/mux"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde una funcion anonima")
}

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Usuario "+this.name)
}

func main() {
	user := &User{"Theo"}
	mux := mux.CreateMux()
	mux.AddFunc("/hola", Hola)
	mux.AddHandle("/usuario", user)
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
