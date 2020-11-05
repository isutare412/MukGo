package common

import "go.mongodb.org/mongo-driver/bson/primitive"

// Restaurant model.
type Restaurant struct {
	ID    primitive.ObjectID
	Name  string
	Coord Coordinate
}

// Review model.
type Review struct {
	ID        primitive.ObjectID
	UserID    string
	UserName  string
	Score     int32
	Comment   string
	Menus     []string
	Wait      bool
	NumPeople int32
	Timestamp int64
}

// User model.
type User struct {
	UserID      string
	Name        string
	Exp         int64
	ReviewCount int32
	LikeCount   int32
}
