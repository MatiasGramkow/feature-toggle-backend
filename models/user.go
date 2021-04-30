package models

// User Model
type User struct {
	ID        uint   `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}
