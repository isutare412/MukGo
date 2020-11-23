package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// User contains user information.
type User struct {
	UserID      string `bson:"user_id"`
	Name        string `bson:"name"`
	Exp         int64  `bson:"exp"`
	ReviewCount int32  `bson:"review_count"`
	LikeCount   int32  `bson:"like_count"`
}

// Review contains review data from user.
type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	RestID    primitive.ObjectID `bson:"restaurant_id"`
	Score     int32              `bson:"score"`
	Comment   string             `bson:"comment"`
	Menus     []string           `bson:"menus"`
	Wait      bool               `bson:"wait"`
	NumPeople int32              `bson:"num_people"`
	Timestamp int64              `bson:"timestamp"`
}

// Like contains like feature data on each review.
type Like struct {
	ReviewID     primitive.ObjectID `bson:"review_id"`
	LikingUserID string             `bson:"liking_user_id"`
	LikedUserID  string             `bson:"liked_user_id"`
}

// Restaurant contains restaurant information.
type Restaurant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Latitude  float64            `bson:"latitude"`
	Longitude float64            `bson:"longitude"`
	Type      int32              `bson:"type"`
}

// Collection names of MongoDB.
const (
	CNUser       string = "users"
	CNReview     string = "reviews"
	CNRestaurant string = "restaurants"
	CNLike       string = "likes"
)
