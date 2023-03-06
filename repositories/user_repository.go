package repositories

import (
	"fmt"
	"github/jorgemvv01/go-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (models.User, error)
	GetAll() (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *models.User) error {
	return ur.db.Create(&user).Error
}

func (ur *userRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	if err := ur.db.Find(&user, id).Error; err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

func (ur *userRepository) GetAll() (*[]models.User, error) {
	var users *[]models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
