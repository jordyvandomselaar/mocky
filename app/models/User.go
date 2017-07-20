package models

import "github.com/jinzhu/gorm"

// User is a model for the tabel users.
type User struct {
	gorm.Model

	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
