package middlewares

import (
	"net/http"
)

// AdminOnly checks for a static API key in the X-Admin-Key header.
// Set ADMIN_KEY in your .env to something long and random.
func AdminOnly(adminKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Admin-Key") != adminKey {
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
