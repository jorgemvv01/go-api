package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/repositories"
	"github/jorgemvv01/go-api/utils"
	"net/http"
	"time"
)

type RentController interface {
	Create(c *gin.Context)
}

type rentController struct {
	rentRepository repositories.RentRepository
}

func NewRentController(repository repositories.RentRepository) RentController {
	return &rentController{
		rentRepository: repository,
	}
}

// CreateRent
// @Summary Create rent
// @Description Create a new rent.
// @Param tags body models.RentRequest true "Create rent"
// @Produce application/json
// @Tags Rent
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /rent/create [post]
func (rc *rentController) Create(c *gin.Context) {
	var rent *models.RentRequest
	if err := c.ShouldBindJSON(&rent); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", rent.StartDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid start date... ` + err.Error(),
		})
		return
	}
	endDate, err := time.Parse("2006-01-02", rent.EndDate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid end date... ` + err.Error(),
		})
		return
	}
	if startDate.After(endDate) {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "The end date must be greater than the start date",
		})
		return
	}
	var days = endDate.Sub(startDate) / (24 * time.Hour)

	rentResponse, err := rc.rentRepository.Create(rent, int(days))
	if err != nil {
		if errors.Is(err, utils.ErrMovieNotFound) || errors.Is(err, utils.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, models.Response{
				Status:  "Error",
				Message: `Unable to create rent... ` + err.Error(),
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Status:  "Error",
				Message: `Unable to create rent... ` + err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Rent created successfully",
		Data:    rentResponse,
	})
}
