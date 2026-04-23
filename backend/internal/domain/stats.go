package domain

type ReviewStats struct {
	TotalReviews  int         `json:"totalReviews"`
	AverageRating float64     `json:"averageRating"`
	RatingCounts  map[int]int `json:"ratingCounts"` // e.g. {1: 10, 2: 5, 3: 20, 4: 50, 5: 100}
}
