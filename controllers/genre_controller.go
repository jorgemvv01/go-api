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
)

type GenreController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type genreController struct {
	repository repositories.GenreRepository
}

func NewGenreController(genreRepository repositories.GenreRepository) GenreController {
	return &genreController{
		repository: genreRepository,
	}
}

// CreateGenre
// @Summary Create Genre
// @Description Create a new genre.
// @Param tags body models.GenreRequest true "Create genre"
// @Produce application/json
// @Tags Movie Genre
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /genres/create [post]
func (gc *genreController) Create(c *gin.Context) {
	var genre *models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	if err := gc.repository.Create(genre); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to create genre... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Genre created successfully",
		Data:    models.NewGenreResponse(*genre),
	})
}

// GetGenreByID
// @Summary Get Genre by ID
// @Description Get a genre by ID.
// @Param ID path string true "Get genre by ID"
// @Produce application/json
// @Tags Movie Genre
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /genres/{ID} [get]
func (gc *genreController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid genre ID",
		})
		return
	}
	genre, err := gc.repository.GetByID(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get genre... ` + err.Error(),
		})
		return
	}
	if genre.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
			Status:  "Error",
			Message: fmt.Sprintf("Genre with ID %d not found", uint(id)),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Genre found",
		Data:    genre,
	})
}

// GetAllGenres
// @Summary Get all Genres
// @Description Get all Genres.
// @Produce application/json
// @Tags Movie Genre
// @Success 200 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /genres [get]
func (gc *genreController) GetAll(c *gin.Context) {
	genres, err := gc.repository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get genres... ` + err.Error(),
		})
		return
	}
	if len(*genres) == 0 {
		c.JSON(http.StatusOK, models.Response{
			Status:  "Success",
			Message: "No genres found",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Genres found",
		Data:    genres,
	})
}

// UpdateGenre
// @Summary Update Genre
// @Description Update Genre by ID.
// @Produce application/json
// @Param ID path string true "Update genre by ID"
// @Param tags body models.GenreRequest true "Update genre"
// @Tags Movie Genre
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /genres/update/{ID} [put]
func (gc *genreController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid genre ID",
		})
		return
	}
	var genre *models.Genre
	if err = c.ShouldBindJSON(&genre); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	var genreResponse *models.GenreResponse
	if genreResponse, err = gc.repository.Update(uint(id), genre); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: fmt.Sprintf("Genre with ID %d not found", uint(id)),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to update genre... ` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Genre updated successfully",
		Data:    genreResponse,
	})
}

// DeleteGenre
// @Summary Delete Genre
// @Description Delete Genre by ID.
// @Produce application/json
// @Param ID path string true "Delete genre by ID"
// @Tags Movie Genre
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /genres/delete/{ID} [delete]
func (gc *genreController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid genre ID",
		})
		return
	}
	if err = gc.repository.Delete(uint(id)); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: fmt.Sprintf("Genre with ID %d not found", uint(id)),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to delete genre...` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Genre deleted successfully",
	})
}
