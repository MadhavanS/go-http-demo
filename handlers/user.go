package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// var users = make(map[string]User) // in-memory database
var (
	users  = make(map[int]User)
	nextID = 1
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = nextID
	nextID++
	users[newUser.ID] = newUser

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func UserByIDHandler(writer http.ResponseWriter, request *http.Request) {
	idStr := strings.TrimPrefix(request.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, exists := users[id]

	switch request.Method {
	case http.MethodGet:
		if !exists {
			http.Error(writer, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(writer).Encode(user)

	case http.MethodPut:
		if !exists {
			http.Error(writer, "User not found", http.StatusNotFound)
			return
		}
		var updated User
		if err := json.NewDecoder(request.Body).Decode(&updated); err != nil {
			http.Error(writer, "User not found", http.StatusBadRequest)
			return
		}

		updated.ID = id
		users[id] = updated
		json.NewEncoder(writer).Encode(updated)

	case http.MethodDelete:
		if !exists {
			http.Error(writer, "User not found", http.StatusNotFound)
			return
		}
		delete(users, id)
		writer.WriteHeader(http.StatusNoContent)
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

/*
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
*/

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	queryName := r.URL.Query().Get("name")

	var result []User
	for _, user := range users {
		if queryName == "" ||
			strings.Contains(strings.ToLower(user.Name), strings.ToLower(queryName)) {
			result = append(result, user)
		}
	}

	json.NewEncoder(w).Encode(result)
}
