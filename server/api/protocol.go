package api

/******************************************************************************
* OAuth 2.0
******************************************************************************/

// UserClaim contains data received from oauth request.
type UserClaim struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
}
