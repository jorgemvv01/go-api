package repositories

import (
	"fmt"
	"github/jorgemvv01/go-api/models"
	"gorm.io/gorm"
)

type TypeRepository interface {
	Create(typeMovie *models.Type) error
	GetByID(id uint) (*models.Type, error)
	GetAll() (*[]models.Type, error)
	Update(id uint, typeMovie *models.Type) (*models.Type, error)
	Delete(id uint) error
}

type typeRepository struct {
	db *gorm.DB
}

func NewTypeRepository(db *gorm.DB) TypeRepository {
	return &typeRepository{
		db: db,
	}
}

func (tr *typeRepository) Create(typeMovie *models.Type) error {
	return tr.db.Create(&typeMovie).Error
}

func (tr *typeRepository) GetByID(id uint) (*models.Type, error) {
	var typeMovie *models.Type
	if err := tr.db.Find(&typeMovie, id).Error; err != nil {
		return typeMovie, err
	}
	if typeMovie.ID == 0 {
		return typeMovie, fmt.Errorf("type with ID %d not found", id)
	}
	return typeMovie, nil
}

func (tr *typeRepository) GetAll() (*[]models.Type, error) {
	var typesMovie *[]models.Type
	if err := tr.db.Find(&typesMovie).Error; err != nil {
		return typesMovie, err
	}
	return typesMovie, nil
}

func (tr *typeRepository) Update(id uint, typeMovie *models.Type) (*models.Type, error) {
	var oldTypeMovie *models.Type
	if err := tr.db.Find(&oldTypeMovie, id).Error; err != nil {
		return typeMovie, err
	}
	if oldTypeMovie.ID == 0 {
		return typeMovie, fmt.Errorf("type with ID %d not found", id)
	}
	oldTypeMovie.Name = typeMovie.Name
	if err := tr.db.Save(&oldTypeMovie).Error; err != nil {
		return oldTypeMovie, err
	}
	return oldTypeMovie, nil
}

func (tr *typeRepository) Delete(id uint) error {
	var typeMovie *models.Type
	if err := tr.db.Find(&typeMovie, id).Error; err != nil {
		return err
	}
	if typeMovie.ID == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}
	return tr.db.Delete(&typeMovie).Error
}
