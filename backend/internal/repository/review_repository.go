package repository

import (
	"context"
	"fmt"
	"math"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mwkhettab/life-review/backend/internal/domain"
)

type ReviewRepository struct {
	db *pgxpool.Pool
}

func NewReviewRepository(db *pgxpool.Pool) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(ctx context.Context, review *domain.Review) (*domain.Review, error) {
	query := `
		INSERT INTO reviews (name, rating, body, pros, cons, do_it_again)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, rating, body, pros, cons, do_it_again, approved, created_at
	`
	row := r.db.QueryRow(ctx, query,
		review.Name,
		review.Rating,
		review.Body,
		review.Pros,
		review.Cons,
		review.DoItAgain,
	)
	return scanReview(row)
}

// GetApproved returns approved reviews with offset-based pagination.
func (r *ReviewRepository) GetApproved(ctx context.Context, limit, offset int) ([]domain.Review, error) {
	query := `
		SELECT id, name, rating, body, pros, cons, do_it_again, approved, created_at
		FROM reviews
		WHERE approved = TRUE
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get reviews: %w", err)
	}
	defer rows.Close()

	var reviews []domain.Review
	for rows.Next() {
		rv, err := scanReview(rows)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, *rv)
	}
	return reviews, nil
}

func (r *ReviewRepository) Approve(ctx context.Context, id int64) (*domain.Review, error) {
	query := `
		UPDATE reviews SET approved = TRUE
		WHERE id = $1
		RETURNING id, name, rating, body, pros, cons, do_it_again, approved, created_at
	`
	return scanReview(r.db.QueryRow(ctx, query, id))
}

func (r *ReviewRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM reviews WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete review id=%d: %w", id, err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("review id=%d not found", id)
	}
	return nil
}

// ── scanner ───────────────────────────────────────────────────────────────────

type rowScanner interface {
	Scan(dest ...any) error
}

func scanReview(row rowScanner) (*domain.Review, error) {
	var rv domain.Review

	// pgx v5 maps TEXT[] directly to []string
	if err := row.Scan(
		&rv.ID,
		&rv.Name,
		&rv.Rating,
		&rv.Body,
		&rv.Pros,
		&rv.Cons,
		&rv.DoItAgain,
		&rv.Approved,
		&rv.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("scan review: %w", err)
	}

	// Format date for the frontend: "February 11, 2026"
	rv.Date = rv.CreatedAt.Format("January 2, 2006")

	// Never return nil slices — frontend expects []
	if rv.Pros == nil {
		rv.Pros = []string{}
	}
	if rv.Cons == nil {
		rv.Cons = []string{}
	}

	return &rv, nil
}

func (r *ReviewRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Review, error) {
	query := `
		SELECT id, name, rating, body, pros, cons, do_it_again, approved, created_at
		FROM reviews
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get all reviews: %w", err)
	}
	defer rows.Close()

	var reviews []domain.Review
	for rows.Next() {
		rv, err := scanReview(rows)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, *rv)
	}
	return reviews, nil
}

func (r *ReviewRepository) GetStats(ctx context.Context) (domain.ReviewStats, error) {
	query := `
		SELECT rating, COUNT(*) AS count
		FROM reviews
		WHERE approved = TRUE
		GROUP BY rating
		ORDER BY rating
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return domain.ReviewStats{}, fmt.Errorf("get stats: %w", err)
	}
	defer rows.Close()

	stats := domain.ReviewStats{
		RatingCounts: map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
	}

	var totalStars int

	for rows.Next() {
		var rating, count int
		if err := rows.Scan(&rating, &count); err != nil {
			return domain.ReviewStats{}, fmt.Errorf("scan stats row: %w", err)
		}
		stats.RatingCounts[rating] = count
		stats.TotalReviews += count
		totalStars += rating * count
	}

	if err := rows.Err(); err != nil {
		return domain.ReviewStats{}, fmt.Errorf("stats rows error: %w", err)
	}

	if stats.TotalReviews > 0 {
		raw := float64(totalStars) / float64(stats.TotalReviews)
		stats.AverageRating = math.Round(raw*100) / 100
	}

	return stats, nil
}
