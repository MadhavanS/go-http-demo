package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fmt.Printf("➡️  %s %s\n", req.Method, req.URL.Path)

		next.ServeHTTP(writer, req)

		duration := time.Since(start)
		fmt.Printf("✅ %s %s took %v\n", req.Method, req.URL.Path, duration)
	})
}
