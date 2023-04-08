package models

import "gorm.io/gorm"

type MovieRent struct {
	gorm.Model
	RentID  uint
	Rent    Rent `gorm:"foreignKey:RentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
