package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Surname  string `binding:"required" gorm:"not null"`
	Lastname string `binding:"required" gorm:"not null"`
}

type UserRequest struct {
	Surname  string `json:"surname"`
	Lastname string `json:"lastname"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Surname  string `json:"surname"`
	Lastname string `json:"lastname"`
}

func NewUserResponse(user User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Surname:  user.Surname,
		Lastname: user.Lastname,
	}
}
