package creds

// User allows the server to authenticate a user
type User struct {
	Email    string `json:"email",db:"email"`
	Password string `json:"password",db:"password"`
	Duration string `json:"duration"`
}
