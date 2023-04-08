package models

import (
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID    uint    `json:"user_id" binding:"required" gorm:"foreignKey:GenreID"`
	Total     float64 `json:"total" binding:"required" gorm:"not null"`
	StartDate string  `json:"start_date" binding:"required" gorm:"not null"`
	EndDate   string  `json:"end_date" binding:"required" gorm:"not null"`
}

type RentRequest struct {
	UserID    uint   `json:"user_id"`
	MovieIDs  []int  `json:"movie_ids"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type RentResponse struct {
	UserID    uint           `json:"user_id"`
	Total     float64        `json:"total"`
	Movies    []MovieSummary `json:"movies"`
	StartDate string         `json:"start_date"`
	EndDate   string         `json:"end_date"`
}
