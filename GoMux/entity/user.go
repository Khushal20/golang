package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
}

type UserLogin struct {
	Password  string `json:"password"`
	Email     string `json:"email" gorm:"unique"`
}

