package api

// CAUserPost defines post data for user.
type CAUserPost struct {
	UserID int    `json:"userid"`
	Name   string `json:"name"`
}

// CAReviewPost defines post data for review.
type CAReviewPost struct {
	UserID  int    `json:"userid"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// CARestaurantPost defines post data for restaurant.
type CARestaurantPost struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
