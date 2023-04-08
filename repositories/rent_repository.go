package repositories

import (
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/utils"
	"gorm.io/gorm"
)

type RentRepository interface {
	Create(rentRequest *models.RentRequest, days int) (*models.RentResponse, error)
}

type rentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) RentRepository {
	return &rentRepository{
		db: db,
	}
}

func (rr *rentRepository) Create(rentRequest *models.RentRequest, days int) (*models.RentResponse, error) {
	var user *models.User
	if err := rr.db.Find(&user, rentRequest.UserID).Error; err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, utils.ErrUserNotFound
	}
	var movies []models.MovieSummary
	for _, movieID := range rentRequest.MovieIDs {
		var movie *models.Movie
		if err := rr.db.Find(&movie, movieID).Error; err != nil {
			return nil, err
		}
		if movie.ID == 0 {
			return nil, utils.ErrMovieNotFound
		}
		movies = append(movies, *models.NewMovieSummary(*movie))
	}
	var total = utils.CalculateTotalRent(movies, days)

	tx := rr.db.Begin()

	var rent = models.Rent{
		UserID:    rentRequest.UserID,
		Total:     total,
		StartDate: rentRequest.StartDate,
		EndDate:   rentRequest.EndDate,
	}

	if err := tx.Create(&rent).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, movie := range movies {
		var movieRent = models.MovieRent{
			MovieID: movie.ID,
			RentID:  rent.ID,
		}
		if err := tx.Create(&movieRent).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &models.RentResponse{
		UserID:    rent.UserID,
		Total:     total,
		Movies:    movies,
		StartDate: rent.StartDate,
		EndDate:   rent.EndDate,
	}, nil
}
