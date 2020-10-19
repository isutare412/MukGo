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
	UserID   string
	UserName string
	Score    int32
	Comment  string
}
