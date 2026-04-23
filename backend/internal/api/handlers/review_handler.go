package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mwkhettab/life-review/backend/internal/domain"
	"github.com/mwkhettab/life-review/backend/internal/services"
)

type ReviewHandler struct {
	service *services.ReviewService
}

func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

// POST /reviews
func (h *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name      string   `json:"name"`
		Rating    int      `json:"rating"`
		Body      string   `json:"body"`
		Pros      []string `json:"pros"`
		Cons      []string `json:"cons"`
		DoItAgain bool     `json:"doItAgain"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Default empty slices so validation and DB don't see nil
	if body.Pros == nil {
		body.Pros = []string{}
	}
	if body.Cons == nil {
		body.Cons = []string{}
	}

	review, err := h.service.SubmitReview(r.Context(), domain.Review{
		Name:      body.Name,
		Rating:    body.Rating,
		Body:      body.Body,
		Pros:      body.Pros,
		Cons:      body.Cons,
		DoItAgain: body.DoItAgain,
	})
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResponse(w, review, http.StatusCreated)
}

// GET /reviews?limit=10&offset=0
func (h *ReviewHandler) List(w http.ResponseWriter, r *http.Request) {
	limit := queryInt(r, "limit", 10)
	offset := queryInt(r, "offset", 0)

	reviews, err := h.service.GetReviews(r.Context(), limit, offset)
	if err != nil {
		jsonError(w, "failed to fetch reviews", http.StatusInternalServerError)
		return
	}

	if reviews == nil {
		reviews = []domain.Review{}
	}

	jsonResponse(w, reviews, http.StatusOK)
}

// PUT /approve/{id}
func (h *ReviewHandler) Approve(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		jsonError(w, "invalid id", http.StatusBadRequest)
		return
	}

	review, err := h.service.ApproveReview(r.Context(), id)
	if err != nil {
		jsonError(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonResponse(w, review, http.StatusOK)
}

// DELETE /reviews/{id}
func (h *ReviewHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		jsonError(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteReview(r.Context(), id); err != nil {
		jsonError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ── helpers ───────────────────────────────────────────────────────────────────

func jsonResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func queryInt(r *http.Request, key string, fallback int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return fallback
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return n
}

// GET /admin/reviews?limit=10&offset=0
func (h *ReviewHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	limit := queryInt(r, "limit", 10)
	offset := queryInt(r, "offset", 0)

	reviews, err := h.service.GetAllReviews(r.Context(), limit, offset)
	if err != nil {
		jsonError(w, "failed to fetch reviews", http.StatusInternalServerError)
		return
	}

	if reviews == nil {
		reviews = []domain.Review{}
	}

	jsonResponse(w, reviews, http.StatusOK)
}

// GET /reviews/stats
func (h *ReviewHandler) Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetStats(r.Context())
	if err != nil {
		jsonError(w, "failed to fetch stats", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, stats, http.StatusOK)
}
