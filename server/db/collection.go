package db

// User contains user information.
type User struct {
	UserID int
	Name   string
}

// Review contains review data from user.
type Review struct {
	UserID  int
	Score   int
	Comment string
}

// Collection names of MongoDB.
const (
	CNUser   string = "users"
	CNReview string = "reviews"
)
