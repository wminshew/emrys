package creds

// Miner allows the server to authenticate a user
type Miner struct {
	Email    string `json:"email",db:"email"`
	Password string `json:"password",db:"password"`
	Duration string `json:"duration"`
}
