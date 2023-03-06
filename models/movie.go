package models

import "time"

type Movie struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Overview    string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	TypeID      uint
	Type        Type `gorm:"foreignKey:TypeID"`
	GenreID     uint
	Genre       Genre     `gorm:"foreignKey:GenreID"`
	ReleaseDate time.Time `gorm:"not null"`
}
