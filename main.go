package main

import (
	"fmt"
	"go-http-demo/handlers"
	"go-http-demo/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateUserHandler(w, r)
			return
		} else if r.Method == http.MethodGet {
			handlers.ListUsersHandler(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	loggedMux := middleware.LoggingMiddleware(mux)

	// TODO: start server on :8081
	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", loggedMux)
}
