package db

import "go.mongodb.org/mongo-driver/bson/primitive"

// User contains user information.
type User struct {
	UserID string `bson:"user_id"`
	Name   string `bson:"name"`
	Exp    int64  `bson:"exp"`
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
