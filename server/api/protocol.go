package api

/******************************************************************************
* models
******************************************************************************/

// Restaurant model for JSON marshaling.
type Restaurant struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/******************************************************************************
* Client to API
******************************************************************************/

// CAUserPost defines post data for user.
type CAUserPost struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

// CAUserGet contains data request for user.
type CAUserGet struct {
	UserID string `json:"user_id"`
}

// CAReviewPost defines post data for review.
type CAReviewPost struct {
	UserID  string `json:"user_id"`
	RestID  string `json:"restaurant_id"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// CARestaurantPost defines post data for restaurant.
type CARestaurantPost struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// CARestaurantsPost defines post data for restaurants.
type CARestaurantsPost struct {
	Restaurants []*Restaurant `json:"restaurants"`
}

// CARestaurantsGet request restarants data within user's sight.
type CARestaurantsGet struct {
	UserID    string  `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/******************************************************************************
* API to Client
******************************************************************************/

// ACUserInfo contains user data.
type ACUserInfo struct {
	Name        string  `json:"name"`
	Level       int     `json:"level"`
	TotalExp    int64   `json:"total_exp"`
	LevelExp    int64   `json:"level_exp"`
	CurExp      int64   `json:"cur_exp"`
	ExpRatio    float64 `json:"exp_ratio"`
	SightRadius float64 `json:"sight_radius"`
}

// ACRestaurantsInfo contains multiple restaurants data.
type ACRestaurantsInfo struct {
	Restaurants []*Restaurant `json:"restaurants"`
}
