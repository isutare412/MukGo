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

// Review model for JSON marshaling.
type Review struct {
	UserName string `json:"user_name"`
	Score    int    `json:"score"`
	Comment  string `json:"comment"`
}

/******************************************************************************
* Client to API
******************************************************************************/

// CAReviewsGet requests reviews of restaurant with id.
type CAReviewsGet struct {
	RestID string `json:"restaurant_id"`
}

// CAReviewPost defines post data for review.
type CAReviewPost struct {
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
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/******************************************************************************
* OAuth 2.0
******************************************************************************/

// UserClaim contains data received from oauth request.
type UserClaim struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
}

/******************************************************************************
* API to Client
******************************************************************************/

// ACErrorReason contains error code. Error should be defined in code package
// from protobuf.
type ACErrorReason struct {
	Code int32 `json:"code"`
}

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

// ACReviewsInfo contains multiple restaurants data.
type ACReviewsInfo struct {
	Reviews []*Review `json:"reviews"`
}
