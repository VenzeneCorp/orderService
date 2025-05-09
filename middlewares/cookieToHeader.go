package middlewares

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		refreshCookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "Refresh token not found", http.StatusUnauthorized)
			return
		}
		refreshToken := refreshCookie.Value

		// Extract token from cookie
		accessCookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Access token not found", http.StatusUnauthorized)
			return
		}

		accessToken := accessCookie.Value

		// Parse token
		token, err := jwt.ParseWithClaims(accessToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID := claims.ID
		role := claims.Subject

		if userID == "" || role == "" {
			http.Error(w, "Missing userID or role in token", http.StatusUnauthorized)
			return
		}

		// Set headers
		r.Header.Set("X-ID", userID)
		r.Header.Set("X-Role", role)
		r.Header.Set("X-Access-Token", accessToken)
		r.Header.Set("X-Refresh-Token", refreshToken)

		next.ServeHTTP(w, r)
	})
}
