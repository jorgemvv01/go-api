package repositories

import (
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/utils"
	"gorm.io/gorm"
)

type GenreRepository interface {
	Create(genre *models.Genre) error
	GetByID(id uint) (*models.Genre, error)
	GetAll() (*[]models.Genre, error)
	Update(id uint, genre *models.Genre) (*models.Genre, error)
	Delete(id uint) error
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{
		db: db,
	}
}

func (gr *genreRepository) Create(genre *models.Genre) error {
	return gr.db.Create(&genre).Error
}

func (gr *genreRepository) GetByID(id uint) (*models.Genre, error) {
	var genre *models.Genre
	if err := gr.db.Find(&genre, id).Error; err != nil {
		return nil, err
	}
	if genre.ID == 0 {
		return nil, utils.ErrNotFound
	}
	return genre, nil
}

func (gr *genreRepository) GetAll() (*[]models.Genre, error) {
	var genres *[]models.Genre
	if err := gr.db.Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (gr *genreRepository) Update(id uint, genre *models.Genre) (*models.Genre, error) {
	var oldGenre *models.Genre
	if err := gr.db.Find(&oldGenre, id).Error; err != nil {
		return nil, err
	}
	if oldGenre.ID == 0 {
		return nil, utils.ErrNotFound
	}
	oldGenre.Name = genre.Name
	if err := gr.db.Save(&oldGenre).Error; err != nil {
		return nil, err
	}
	return oldGenre, nil
}

func (gr *genreRepository) Delete(id uint) error {
	var genre *models.Genre
	if err := gr.db.Find(&genre, id).Error; err != nil {
		return err
	}
	if genre.ID == 0 {
		return utils.ErrNotFound
	}
	return gr.db.Delete(&genre).Error
}
