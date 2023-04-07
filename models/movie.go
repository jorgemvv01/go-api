package models

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	Name        string  `json:"name" binding:"required" gorm:"not null"`
	Overview    string  `json:"overview" binding:"required" gorm:"not null"`
	Price       float64 `json:"price" binding:"required" gorm:"not null"`
	TypeID      uint    `json:"type_id" binding:"required" gorm:"foreignKey:GenreID"`
	GenreID     uint    `json:"genre_id" binding:"required" gorm:"foreignKey:GenreID"`
	ReleaseDate string  `json:"release_date" binding:"required" gorm:"not null"`
}

type MovieRequest struct {
	Name        string    `json:"name"`
	Overview    string    `json:"overview"`
	Price       float64   `json:"price"`
	TypeID      uint      `json:"type_id"`
	GenreID     uint      `json:"genre_id"`
	ReleaseDate time.Time `json:"release_date"`
}

type MovieResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Overview    string        `json:"overview"`
	Price       float64       `json:"price"`
	Type        TypeResponse  `json:"type"`
	Genre       GenreResponse `json:"genre"`
	ReleaseDate string        `json:"release_date"`
}

func NewMovieResponse(movie Movie, movieType Type, movieGenre Genre) *MovieResponse {
	return &MovieResponse{
		ID:          movie.ID,
		Name:        movie.Name,
		Overview:    movie.Overview,
		Price:       movie.Price,
		Type:        *NewTypeResponse(movieType),
		Genre:       *NewGenreResponse(movieGenre),
		ReleaseDate: movie.ReleaseDate,
	}
}
