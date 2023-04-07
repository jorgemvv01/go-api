package repositories

import (
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.UserResponse, error)
	GetAll() (*[]models.UserResponse, error)
	Update(id uint, user *models.User) (*models.UserResponse, error)
	Delete(id uint) error
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

func (ur *userRepository) GetByID(id uint) (*models.UserResponse, error) {
	var user *models.User
	if err := ur.db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return models.NewUserResponse(*user), nil
}

func (ur *userRepository) GetAll() (*[]models.UserResponse, error) {
	var users *[]models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	var usersResponse []models.UserResponse
	for _, user := range *users {
		usersResponse = append(usersResponse, *models.NewUserResponse(user))
	}
	return &usersResponse, nil
}

func (ur *userRepository) Update(id uint, user *models.User) (*models.UserResponse, error) {
	var oldUser *models.User
	if err := ur.db.Find(&oldUser, id).Error; err != nil {
		return nil, err
	}
	if oldUser.ID == 0 {
		return nil, utils.ErrNotFound
	}
	oldUser.Surname = user.Surname
	oldUser.Lastname = user.Lastname
	if err := ur.db.Save(&oldUser).Error; err != nil {
		return nil, err
	}
	return models.NewUserResponse(*oldUser), nil
}

func (ur *userRepository) Delete(id uint) error {
	var user *models.User
	if err := ur.db.Find(&user, id).Error; err != nil {
		return err
	}
	if user.ID == 0 {
		return utils.ErrNotFound
	}
	return ur.db.Delete(&user).Error
}
