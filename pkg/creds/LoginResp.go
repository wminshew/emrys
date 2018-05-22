package creds

// LoginResp holds a valid JWT after successful login
type LoginResp struct {
	Token string `json:"token"`
}
