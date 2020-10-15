package api

/******************************************************************************
* models
******************************************************************************/

// Restaurant model for JSON marshaling.
type Restaurant struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/******************************************************************************
* Client to API
******************************************************************************/

// CAUserPost defines post data for user.
type CAUserPost struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
}

// CAUserGet contains data request for user.
type CAUserGet struct {
	UserID string `json:"userid"`
}

// CAReviewPost defines post data for review.
type CAReviewPost struct {
	UserID  string `json:"userid"`
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
	UserID    string  `json:"userid"`
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
	TotalExp    int64   `json:"totalExp"`
	LevelExp    int64   `json:"levelExp"`
	ExpRatio    float64 `json:"expRatio"`
	SightRadius float64 `json:"sightRadius"`
}

// ACRestaurantsInfo contains multiple restaurants data.
type ACRestaurantsInfo struct {
	Restaurants []*Restaurant `json:"restaurants"`
}
