package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mwkhettab/life-review/backend/internal/api/handlers"
	"github.com/mwkhettab/life-review/backend/internal/repository"
	"github.com/mwkhettab/life-review/backend/internal/services"
)

type Container struct {
	ReviewHandler *handlers.ReviewHandler
}

func NewContainer(db *pgxpool.Pool) *Container {
	reviewRepo := repository.NewReviewRepository(db)
	reviewSvc := services.NewReviewService(reviewRepo)
	reviewHandler := handlers.NewReviewHandler(reviewSvc)

	return &Container{
		ReviewHandler: reviewHandler,
	}
}
