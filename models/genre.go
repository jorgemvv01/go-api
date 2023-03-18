package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type GenreRequest struct {
	Name string `json:"name"`
}
