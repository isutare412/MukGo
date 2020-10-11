package api

// JSONUserPost defines post data for user.
type JSONUserPost struct {
	UserID int    `json:"userid"`
	Name   string `json:"name"`
}

// JSONReviewPost defines post data for review.
type JSONReviewPost struct {
	UserID  int    `json:"userid"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// JSONRestaurantPost defines post data for restaurant.
type JSONRestaurantPost struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}
