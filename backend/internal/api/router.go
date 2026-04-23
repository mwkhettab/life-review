package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/mwkhettab/life-review/backend/internal/api/handlers"
	"github.com/mwkhettab/life-review/backend/internal/api/middlewares"
)

func NewRouter(reviewHandler *handlers.ReviewHandler, adminKey string, allowedOrigins []string) http.Handler {
	r := chi.NewRouter()

	// ── Global middleware ────────────────────────────────────────────────────

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP) // trust X-Forwarded-For so rate limiting uses real IP
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Security headers on every response
	r.Use(middlewares.SecureHeaders)

	// Cap request body at 64KB — more than enough for a review
	r.Use(middleware.RequestSize(64 * 1024))

	// CORS — only allow your frontend origin(s)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins, // e.g. ["https://yoursite.com"]
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-Admin-Key"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// ── Rate limiting ────────────────────────────────────────────────────────
	//
	// Global: 100 requests per minute per IP — protects all routes.
	// POST /reviews gets a tighter limit (5/min) to prevent review spam.
	//
	// In-memory is fine here; at 5k users you'd need ~100 concurrent
	// hammering clients to even approach any meaningful threshold.

	r.Use(httprate.LimitByIP(100, time.Minute))

	// ── Public routes ────────────────────────────────────────────────────────

	r.Get("/reviews", reviewHandler.List)
	r.Get("/stats", reviewHandler.Stats)

	// Tighter rate limit on submission: 5 per IP per minute
	r.With(httprate.LimitByIP(5, time.Minute)).Post("/reviews", reviewHandler.Create)

	// ── Admin-only routes ────────────────────────────────────────────────────

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AdminOnly(adminKey))
		r.Get("/admin/reviews", reviewHandler.ListAll)
		r.Put("/approve/{id}", reviewHandler.Approve)
		r.Delete("/reviews/{id}", reviewHandler.Delete)
	})

	return r
}
