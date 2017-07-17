package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Comment struct {
	Title   string
	Content template.HTML
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var error = template.Must(template.ParseFiles("states_page/error.html"))

func handlerError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	error.Execute(w, nil)
}

func renderTemplates(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		// http.Error(w, "No es posible retornar el template.", http.StatusInternalServerError)
		handlerError(w, http.StatusInternalServerError)
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	title := "Ejemplo."
	content := template.HTML("Este es el contenido del comentario <strong>Test</strong>.")
	comment := Comment{
		Title:   title,
		Content: content,
	}
	renderTemplates(w, r, "comments/comment", comment)
}

func main() {
	staticFiles := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.HandleFunc("/comment/", comment)

	mux.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	fmt.Println("Servidor a la escucha en el puerto :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
