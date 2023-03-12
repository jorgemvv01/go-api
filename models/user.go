package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Surname  string `json:"Surname" binding:"required" gorm:"not null"`
	Lastname string `json:"Lastname" binding:"required" gorm:"not null"`
}

type UserRequest struct {
	Surname  string `json:"surname"`
	Lastname string `json:"lastname"`
}
