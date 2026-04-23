package domain

import "time"

type Review struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Rating    int       `json:"rating"` // 1–5
	Body      string    `json:"body"`
	Pros      []string  `json:"pros"`
	Cons      []string  `json:"cons"`
	DoItAgain bool      `json:"doItAgain"`
	Approved  bool      `json:"approved"`
	Date      string    `json:"date"`      // "February 11, 2026"
	CreatedAt time.Time `json:"-"`
}

// Limits enforced in the service layer
const (
	MaxNameLen    = 100
	MaxBodyLen    = 2000
	MaxTagLen     = 150 // individual pro or con
	MaxTags       = 6
)
