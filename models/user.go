package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Surname  string `json:"surname" binding:"required" gorm:"not null"`
	Lastname string `json:"lastname" binding:"required" gorm:"not null"`
}

type UserRequest struct {
	Surname  string `json:"surname"`
	Lastname string `json:"lastname"`
}
