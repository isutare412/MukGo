package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Restaurant model.
type Restaurant struct {
	ID             primitive.ObjectID
	Name           string
	Coord          Coordinate
	RestaurantType int32
}

// Review model.
type Review struct {
	ID        primitive.ObjectID
	UserID    string
	UserName  string
	UserExp   int64
	Score     int32
	Comment   string
	Menus     []string
	Wait      bool
	NumPeople int32
	Timestamp int64
	LikeCount int32
	LikedByMe bool
}

// User model.
type User struct {
	UserID      string
	Name        string
	Exp         int64
	ReviewCount int32
	LikeCount   int32
	RTCounts    map[int32]int32
}
