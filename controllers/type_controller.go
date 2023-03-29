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

type TypeController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type typeController struct {
	typeRepository repositories.TypeRepository
}

func NewTypeController(repository repositories.TypeRepository) TypeController {
	return &typeController{
		typeRepository: repository,
	}
}

func (tc *typeController) Create(c *gin.Context) {
	var typeMovie *models.Type
	if err := c.ShouldBindJSON(&typeMovie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	if err := tc.typeRepository.Create(typeMovie); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to create type... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Type created successfully",
		Data:    typeMovie,
	})
}

// CreateType
// @Summary Get Type by ID
// @Description Get Type by ID
// @Param ID path string true "Get Type by ID"
// @Produce application/json
// @Tags Movie Type
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 404 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /types/{ID} [get]
func (tc *typeController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid type ID",
		})
		return
	}
	var movieType *models.Type
	if movieType, err = tc.typeRepository.GetByID(uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Success",
			Message: `Unable to get type... ` + err.Error(),
		})
		return
	}
	if movieType.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
			Status:  "Error",
			Message: fmt.Sprintf("Type with ID %d not found", uint(id)),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Type created successfully",
		Data:    movieType,
	})
}

// CreateType
// @Summary Get all Types
// @Description Get all Types
// @Produce application/json
// @Tags Movie Type
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /types/ [get]
func (tc *typeController) GetAll(c *gin.Context) {
	typesMovie, err := tc.typeRepository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get types... ` + err.Error(),
		})
		return
	}
	if len(*typesMovie) == 0 {
		c.JSON(http.StatusOK, models.Response{
			Status:  "Success",
			Message: "No types found",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Types found",
		Data:    typesMovie,
	})
}

func (tc *typeController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid type ID",
		})
		return
	}
	var movieType *models.Type
	if err = c.ShouldBindJSON(&movieType); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	if movieType, err = tc.typeRepository.Update(uint(id), movieType); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: fmt.Sprintf("Type with ID %d not found", uint(id)),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to update type... ` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Type updated successfully",
		Data:    movieType,
	})
}

func (tc *typeController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid type ID",
		})
		return
	}
	if err = tc.typeRepository.Delete(uint(id)); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: fmt.Sprintf("Type with ID %d not found", uint(id)),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to delete type... ` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Type deleted successfully",
	})
}
