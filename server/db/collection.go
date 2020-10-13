package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// User contains user information.
type User struct {
	UserID int    `bson:"userid"`
	Name   string `bson:"name"`
	Exp    int64  `bson:"exp"`
}

// Review contains review data from user.
type Review struct {
	UserID  int    `bson:"userid"`
	Score   int    `bson:"score"`
	Comment string `bson:"comment"`
}

// Restaurant contains restaurant information.
type Restaurant struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Latitude  float64            `bson:"latitude"`
	Longitude float64            `bson:"longitude"`
}

// Collection names of MongoDB.
const (
	CNUser       string = "users"
	CNReview     string = "reviews"
	CNRestaurant string = "restaurants"
)
