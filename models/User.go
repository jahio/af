package models

import "gorm.io/gorm"

// User - a basic user object
type User struct {
	gorm.Model
	Email    string
	Password string
}
