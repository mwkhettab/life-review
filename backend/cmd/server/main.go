package main

import (
	"context"
	"log"

	"github.com/mwkhettab/life-review/backend/config"
	"github.com/mwkhettab/life-review/backend/internal/api"
	"github.com/mwkhettab/life-review/backend/internal/app"
	"github.com/mwkhettab/life-review/backend/internal/database"
	"github.com/mwkhettab/life-review/backend/internal/server"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	container := app.NewContainer(db)
	router := api.NewRouter(container.ReviewHandler, cfg.AdminKey, cfg.AllowedOrigins)
	srv := server.New(cfg.Port, router)

	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
