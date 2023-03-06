package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Surname  string `json:"surname" binding:"required" gorm:"not null"`
	Lastname string `json:"lastname" binding:"required" gorm:"not null"`
}
