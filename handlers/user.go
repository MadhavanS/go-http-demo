package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = make(map[string]User) // in-memory database

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// example url: /user/123
	path := r.URL.Path                // "/user/123"
	parts := strings.Split(path, "/") // ["", "user", "123"]

	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	userID := parts[2]
	switch r.Method {
	case http.MethodGet:
		user, ok := users[userID]
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)

	case http.MethodPost:
		var usr User
		if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		usr.ID = userID
		users[userID] = usr
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(usr)

	case http.MethodPut:
		var usr User
		if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		usr.ID = userID
		users[userID] = usr
		json.NewEncoder(w).Encode(usr)

	case http.MethodDelete:
		delete(users, userID)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
