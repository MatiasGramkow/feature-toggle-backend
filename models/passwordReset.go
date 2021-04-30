package models

// PasswordReset Model
type PasswordReset struct {
	ID    uint
	Email string
	Token string
}
