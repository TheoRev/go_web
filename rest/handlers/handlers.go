package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TheoRev/go_web/rest/models"
	"github.com/gorilla/mux"
)

// GetUsers Se obtiene todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtiene todos los usuarios!")
}

// GetUser Se obtiene un usuario
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	response := models.GetDefaultResponse()
	user, err := models.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFound(err.Error)
	} else {
		response.Data = user
	}

	output, _ := json.Marshal(&response)

	// fmt.Fprintf(w, "Se obtiene un usuario!")
	fmt.Fprintf(w, string(output))
}

// CreateUser Se crea un usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario!")
}

// UpdateUser Se actualiza un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario!")
}

// DeleteUser Se elimina un usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario!")
}
