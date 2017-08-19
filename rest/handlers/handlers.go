package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TheoRev/go_web/rest/models"
	"github.com/gorilla/mux"
)

// GetUsers Se obtiene todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

// GetUser Se obtiene un usuario
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

// CreateUser Se crea un usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	// decoder := json.NewDecoder(r.Body)

	fmt.Fprintf(w, user.Username)

	// if err := decoder.Decode(&user); err != nil {
	// 	models.SendUnprocessableEntity(w)
	// } else {
	// 	models.SendData(w, models.SaveUser(user))
	// }
}

// UpdateUser Se actualiza un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario!")
}

// DeleteUser Se elimina un usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario!")
}
