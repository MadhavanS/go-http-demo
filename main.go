package main

import (
	"fmt"
	"go-http-demo/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", methodHandler(http.MethodGet, handlers.HelloHandler))
	http.HandleFunc("/about", methodHandler(http.MethodGet, handlers.AboutHandler))
	http.HandleFunc("/wish", methodHandler(http.MethodGet, handlers.WishHandler))
	http.HandleFunc("/greet", methodHandler(http.MethodPost, handlers.GreetHandler))
	http.HandleFunc("/add", methodHandler(http.MethodPost, handlers.AddHandler))
	http.HandleFunc("/user/", handlers.UserHandler)

	// TODO: start server on :8081
	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

func methodHandler(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method %s is not allowed", r.Method)
			return
		}
		handler(w, r)
	}
}
