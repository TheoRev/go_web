package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Curso struct {
	Nombre   string
	Duracion int
}

type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Tags     []string
	Cursos   []Curso
}

// var templates = template.Must(template.New("T").ParseFiles(
// 	"templates/index.html",
// 	"templates/footer.html",
// 	"templates/header.html",
// ))

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

func main() {
	tags := []string{"Go", "Java", "PHP", "Ruby", "Python"}
	cursos := []Curso{Curso{"PHP", 12}, Curso{"Java", 60}, Curso{"Go", 10}}
	usuario := Usuario{
		UserName: "Theo",
		Edad:     25,
		Activo:   true,
		Admin:    false,
		Tags:     tags,
		Cursos:   cursos,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplates(w, r, "index", nil)
	})

	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		renderTemplates(w, r, "usuario", usuario)
	})

	// mux := http.NewServeMux()
	// mux.HandleFunc("/comment/", comment)

	fmt.Println("Servidor a la escucha en el puerto :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
