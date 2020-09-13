package api

// RestRequest describes json request from the client.
type RestRequest struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

// RestResponse describes json response to the client.
type RestResponse struct {
	Message string `json:"message"`
}
