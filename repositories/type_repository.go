package repositories

import (
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/utils"
	"gorm.io/gorm"
)

type TypeRepository interface {
	Create(movieType *models.Type) error
	GetByID(id uint) (*models.TypeResponse, error)
	GetAll() (*[]models.TypeResponse, error)
	Update(id uint, movieType *models.Type) (*models.TypeResponse, error)
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

func (tr *typeRepository) Create(movieType *models.Type) error {
	return tr.db.Create(&movieType).Error
}

func (tr *typeRepository) GetByID(id uint) (*models.TypeResponse, error) {
	var movieType *models.Type
	if err := tr.db.Find(&movieType, id).Error; err != nil {
		return nil, err
	}
	return models.NewTypeResponse(*movieType), nil
}

func (tr *typeRepository) GetAll() (*[]models.TypeResponse, error) {
	var movieTypes *[]models.Type
	if err := tr.db.Find(&movieTypes).Error; err != nil {
		return nil, err
	}
	var movieTypeResponse []models.TypeResponse
	for _, movieType := range *movieTypes {
		movieTypeResponse = append(movieTypeResponse, *models.NewTypeResponse(movieType))
	}
	return &movieTypeResponse, nil
}

func (tr *typeRepository) Update(id uint, movieType *models.Type) (*models.TypeResponse, error) {
	var oldMovieType *models.Type
	if err := tr.db.Find(&oldMovieType, id).Error; err != nil {
		return nil, err
	}
	if oldMovieType.ID == 0 {
		return nil, utils.ErrNotFound
	}
	oldMovieType.Name = movieType.Name
	if err := tr.db.Save(&oldMovieType).Error; err != nil {
		return nil, err
	}
	return models.NewTypeResponse(*oldMovieType), nil
}

func (tr *typeRepository) Delete(id uint) error {
	var movieType *models.Type
	if err := tr.db.Find(&movieType, id).Error; err != nil {
		return err
	}
	if movieType.ID == 0 {
		return utils.ErrNotFound
	}
	return tr.db.Delete(&movieType).Error
}
