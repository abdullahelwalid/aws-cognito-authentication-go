package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	FirstName string
	Email string `gorm:"unique"`
	LastName string
}
