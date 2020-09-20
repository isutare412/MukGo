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

// JSONReview defines protocol for review request from a client.
type JSONReview struct {
	UserID  int    `json:"userid"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}
