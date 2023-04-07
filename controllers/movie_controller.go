package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/utils"
	"net/http"
	"strconv"
	"time"
)

type MovieController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type movieController struct {
	movieRepository repositories.MovieRepository
}

func NewMovieController(repository repositories.MovieRepository) MovieController {
	return &movieController{
		movieRepository: repository,
	}
}

// CreateMovie
// @Summary Create Movie
// @Description Create a new movie.
// @Param tags body models.MovieRequest true "Create movie"
// @Produce application/json
// @Tags Movies
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /movies/create [post]
func (mc *movieController) Create(c *gin.Context) {
	var movie *models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", movie.ReleaseDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid release date... ` + err.Error(),
		})
		return
	}
	movie.ReleaseDate = releaseDate.Format("2006-01-02")

	movieResponse, err := mc.movieRepository.Create(movie)
	if err != nil {
		if errors.Is(err, utils.ErrGenreNotFound) || errors.Is(err, utils.ErrTypeNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: `Unable to create movie... ` + err.Error(),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to create movie... ` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Movie created successfully",
		Data:    movieResponse,
	})
}

// GetMovieByID
// @Summary Get Movie by ID
// @Description Get a movie by ID.
// @Param ID path string true "Get movie by ID"
// @Produce application/json
// @Tags Movies
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /movies/{ID} [get]
func (mc *movieController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid movie ID",
		})
		return
	}
	movie, err := mc.movieRepository.GetByID(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get movie... ` + err.Error(),
		})
		return
	}
	if movie.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
			Status:  "Error",
			Message: fmt.Sprintf("Movie with ID %d not found", uint(id)),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Movie found",
		Data:    movie,
	})
}

// GetAllMovies
// @Summary Get all Movies
// @Description Get all Movies.
// @Produce application/json
// @Tags Movies
// @Success 200 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /movies [get]
func (mc *movieController) GetAll(c *gin.Context) {
	movies, err := mc.movieRepository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get movies... ` + err.Error(),
		})
		return
	}
	if len(*movies) == 0 {
		c.JSON(http.StatusOK, models.Response{
			Status:  "Success",
			Message: "No movies found",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Movies found",
		Data:    movies,
	})
}

// UpdateMovie
// @Summary Update Movie
// @Description Update Movie by ID.
// @Produce application/json
// @Param ID path string true "Update movie by ID"
// @Param tags body models.MovieRequest true "Update movie"
// @Tags Movies
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /movies/update/{ID} [put]
func (mc *movieController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid movie ID",
		})
		return
	}
	var movie *models.Movie
	if err = c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	var movieResponse *models.MovieResponse
	if movieResponse, err = mc.movieRepository.Update(uint(id), movie); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: fmt.Sprintf("Movie with ID %d not found", uint(id)),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to update movie... ` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Movie updated successfully",
		Data:    movieResponse,
	})
}

// DeleteMovie
// @Summary Delete Movie
// @Description Delete Movie by ID.
// @Produce application/json
// @Param ID path string true "Delete Movie by ID"
// @Tags Movies
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /movies/delete/{ID} [delete]
func (mc *movieController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid movie ID",
		})
		return
	}
	if err = mc.movieRepository.Delete(uint(id)); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: fmt.Sprintf("Movie with ID %d not found", uint(id)),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to delete movie...` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Movie deleted successfully",
	})
}
