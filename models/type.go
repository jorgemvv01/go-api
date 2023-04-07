package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"not null"`
}

type TypeRequest struct {
	Name string `json:"name"`
}

type TypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewTypeResponse(typeMovie Type) *TypeResponse {
	return &TypeResponse{
		ID:   typeMovie.ID,
		Name: typeMovie.Name,
	}
}
