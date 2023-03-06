package controllers

import (
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/repositories"
	"net/http"
	"strconv"
)

type UserController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type userController struct {
	userRepository repositories.UserRepository
}

func NewUserController(repository repositories.UserRepository) UserController {
	return &userController{
		userRepository: repository,
	}
}

func (uc *userController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": `Invalid request body... ` + err.Error(),
		})
		return
	}
	if err := uc.userRepository.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": `Unable to create user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User created successfully",
		"user":    user,
	})
}

func (ur *userController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid user ID",
		})
		return
	}
	user, err := ur.userRepository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": `Unable to get user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User found",
		"user":    user,
	})
}

func (ur *userController) GetAll(c *gin.Context) {
	users, err := ur.userRepository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": `Unable to get users... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Users found",
		"user":    users,
	})
}
