package models

import "gorm.io/gorm"

type MovieRent struct {
	gorm.Model
	RentID  uint `gorm:"foreignKey:RentID"`
	MovieID uint `gorm:"foreignKey:MovieID"`
}
