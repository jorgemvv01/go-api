package models

import "time"

type Rent struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User      `gorm:"foreignKey:UserID"`
	Total     float64   `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}
