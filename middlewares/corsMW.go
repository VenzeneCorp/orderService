package middlewares

import (
	"net/http"
)

// CORS Middleware
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow CORS from specific origin(s) or "*" for all origins
		w.Header().Set("Access-Control-Allow-Origin", "*") // Or specify specific domain like "http://example.com"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// If it's an OPTIONS request, return early to avoid further processing
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Continue to the next handler
		next.ServeHTTP(w, r)
	})
}
