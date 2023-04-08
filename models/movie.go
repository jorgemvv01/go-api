package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string  `json:"name" binding:"required" gorm:"not null"`
	Overview    string  `json:"overview" binding:"required" gorm:"not null"`
	Price       float64 `json:"price" binding:"required" gorm:"not null"`
	TypeID      uint    `json:"type_id" binding:"required"`
	Type        Type    `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" binding:"-"`
	GenreID     uint    `json:"genre_id" binding:"required"`
	Genre       Genre   `gorm:"foreignKey:GenreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" binding:"-"`
	ReleaseDate string  `json:"release_date" binding:"required" gorm:"not null"`
}

type MovieRequest struct {
	Name        string  `json:"name"`
	Overview    string  `json:"overview"`
	Price       float64 `json:"price"`
	TypeID      uint    `json:"type_id"`
	GenreID     uint    `json:"genre_id"`
	ReleaseDate string  `json:"release_date"`
}

type MovieSummary struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	TypeID uint    `json:"-"`
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

func NewMovieSummary(movie Movie) *MovieSummary {
	return &MovieSummary{
		ID:     movie.ID,
		Name:   movie.Name,
		Price:  movie.Price,
		TypeID: movie.TypeID,
	}
}
