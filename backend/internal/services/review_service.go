package services

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mwkhettab/life-review/backend/internal/domain"
	"github.com/mwkhettab/life-review/backend/internal/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) SubmitReview(ctx context.Context, req domain.Review) (*domain.Review, error) {
	if err := validate(req); err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, &req)
}

func (s *ReviewService) GetReviews(ctx context.Context, limit, offset int) ([]domain.Review, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetApproved(ctx, limit, offset)
}

func (s *ReviewService) ApproveReview(ctx context.Context, id int64) (*domain.Review, error) {
	return s.repo.Approve(ctx, id)
}

func (s *ReviewService) DeleteReview(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

// ── validation ────────────────────────────────────────────────────────────────

func validate(r domain.Review) error {
	// Required
	if strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if utf8.RuneCountInString(r.Name) > domain.MaxNameLen {
		return fmt.Errorf("name must be %d characters or fewer", domain.MaxNameLen)
	}
	if r.Rating < 1 || r.Rating > 5 {
		return fmt.Errorf("rating must be between 1 and 5")
	}

	// Optional — only validate if provided
	if r.Body != "" && utf8.RuneCountInString(r.Body) > domain.MaxBodyLen {
		return fmt.Errorf("body must be %d characters or fewer", domain.MaxBodyLen)
	}

	if len(r.Pros) > domain.MaxTags {
		return fmt.Errorf("pros must have %d items or fewer", domain.MaxTags)
	}
	for _, p := range r.Pros {
		if strings.TrimSpace(p) == "" {
			return fmt.Errorf("pros cannot contain blank items")
		}
		if utf8.RuneCountInString(p) > domain.MaxTagLen {
			return fmt.Errorf("each pro must be %d characters or fewer", domain.MaxTagLen)
		}
	}

	if len(r.Cons) > domain.MaxTags {
		return fmt.Errorf("cons must have %d items or fewer", domain.MaxTags)
	}
	for _, c := range r.Cons {
		if strings.TrimSpace(c) == "" {
			return fmt.Errorf("cons cannot contain blank items")
		}
		if utf8.RuneCountInString(c) > domain.MaxTagLen {
			return fmt.Errorf("each con must be %d characters or fewer", domain.MaxTagLen)
		}
	}

	return nil
}

func (s *ReviewService) GetAllReviews(ctx context.Context, limit, offset int) ([]domain.Review, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *ReviewService) GetStats(ctx context.Context) (domain.ReviewStats, error) {
	return s.repo.GetStats(ctx)
}
