package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"not null"`
}

type GenreRequest struct {
	Name string `json:"name"`
}

type GenreResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewGenreResponse(genre Genre) *GenreResponse {
	return &GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}
}
