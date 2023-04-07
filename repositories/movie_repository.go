package repositories

import (
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/utils"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Create(movie *models.Movie) (*models.MovieResponse, error)
	GetByID(id uint) (*models.MovieResponse, error)
	GetAll() (*[]models.MovieResponse, error)
	Update(id uint, movie *models.Movie) (*models.MovieResponse, error)
	Delete(id uint) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{
		db: db,
	}
}

func (mr *movieRepository) Create(movie *models.Movie) (*models.MovieResponse, error) {
	var movieType models.Type
	if err := mr.db.Find(&movieType, movie.TypeID).Error; err != nil {
		return nil, err
	}
	if movieType.ID == 0 {
		return nil, utils.ErrTypeNotFound
	}
	var movieGenre models.Genre
	if err := mr.db.Find(&movieGenre, movie.GenreID).Error; err != nil {
		return nil, err
	}
	if movieGenre.ID == 0 {
		return nil, utils.ErrGenreNotFound
	}
	if err := mr.db.Create(&movie).Error; err != nil {
		return nil, err
	}
	return models.NewMovieResponse(*movie, movieType, movieGenre), nil
}

func (mr *movieRepository) GetByID(id uint) (*models.MovieResponse, error) {
	var movie *models.Movie
	if err := mr.db.Find(&movie, id).Error; err != nil {
		return nil, err
	}
	var movieType models.Type
	if err := mr.db.Find(&movieType, movie.TypeID).Error; err != nil {
		return nil, err
	}
	var movieGenre models.Genre
	if err := mr.db.Find(&movieGenre, movie.GenreID).Error; err != nil {
		return nil, err
	}
	return models.NewMovieResponse(*movie, movieType, movieGenre), nil
}

func (mr *movieRepository) GetAll() (*[]models.MovieResponse, error) {
	var movies *[]models.Movie
	if err := mr.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	var moviesResponse []models.MovieResponse
	for _, movie := range *movies {
		var movieType models.Type
		if err := mr.db.Find(&movieType, movie.TypeID).Error; err != nil {
			return nil, err
		}
		var movieGenre models.Genre
		if err := mr.db.Find(&movieGenre, movie.GenreID).Error; err != nil {
			return nil, err
		}
		moviesResponse = append(moviesResponse, *models.NewMovieResponse(movie, movieType, movieGenre))
	}
	return &moviesResponse, nil
}

func (mr *movieRepository) Update(id uint, movie *models.Movie) (*models.MovieResponse, error) {
	var oldMovie *models.Movie
	if err := mr.db.Find(&oldMovie, id).Error; err != nil {
		return nil, err
	}
	if oldMovie.ID == 0 {
		return nil, utils.ErrNotFound
	}

	var movieType models.Type
	if err := mr.db.Find(&movieType, movie.TypeID).Error; err != nil {
		return nil, err
	}
	if movieType.ID == 0 {
		return nil, utils.ErrTypeNotFound
	}
	var movieGenre models.Genre
	if err := mr.db.Find(&movieGenre, movie.GenreID).Error; err != nil {
		return nil, err
	}
	if movieGenre.ID == 0 {
		return nil, utils.ErrGenreNotFound
	}

	oldMovie.Name = movie.Name
	oldMovie.Overview = movie.Overview
	oldMovie.Price = movie.Price
	oldMovie.TypeID = movie.TypeID
	oldMovie.GenreID = movie.GenreID
	oldMovie.ReleaseDate = movie.ReleaseDate

	if err := mr.db.Save(&oldMovie).Error; err != nil {
		return nil, err
	}

	return models.NewMovieResponse(*oldMovie, movieType, movieGenre), nil
}

func (mr *movieRepository) Delete(id uint) error {
	var movie *models.Movie
	if err := mr.db.Find(&movie, id).Error; err != nil {
		return err
	}
	if movie.ID == 0 {
		return utils.ErrNotFound
	}
	return mr.db.Delete(&movie).Error
}
