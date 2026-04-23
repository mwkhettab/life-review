package middlewares

import "net/http"

// SecureHeaders sets standard security headers on every response.
func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent browsers from MIME-sniffing the content type
		w.Header().Set("X-Content-Type-Options", "nosniff")
		// Block the response from being embedded in an iframe
		w.Header().Set("X-Frame-Options", "DENY")
		// Basic XSS protection for older browsers
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		// Only send referrer on same origin
		w.Header().Set("Referrer-Policy", "same-origin")
		// Prevent caching of API responses
		w.Header().Set("Cache-Control", "no-store")

		next.ServeHTTP(w, r)
	})
}
