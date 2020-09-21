package api

// JSONRequest describes json request from the client.
type JSONRequest struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

// JSONResponse describes json response to the client.
type JSONResponse struct {
	Message string `json:"message"`
}

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
