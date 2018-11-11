package creds

// Account allows the server to authenticate an account
type Account struct {
	Email    string `json:"email",db:"email"`
	Password string `json:"password",db:"password"`
}
