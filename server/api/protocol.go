package api

// JSONUserAdd defines protocol of information to add new user.
type JSONUserAdd struct {
	UserID int    `json:"userid"`
	Name   string `json:"name"`
}

// JSONReview defines protocol for review request from a client.
type JSONReview struct {
	UserID  int    `json:"userid"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}
