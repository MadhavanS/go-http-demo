package handlers

import (
	"encoding/json"
	"go-http-demo/models"
	"go-http-demo/services"
	"net/http"
)

func ListUsersHandler(writer http.ResponseWriter, request *http.Request) {
	users := services.ListUsers()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(&users)
}

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	services.CreateUser(user)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(user)
}
