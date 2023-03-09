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

// CreateUser
// @Summary Create User
// @Description Create a new user.
// @Param tags body models.UserRequest true "Create user"
// @Produce application/json
// @Tags Users
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Router /users/create [post]
func (uc *userController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	if err := uc.userRepository.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "error",
			Message: `Unable to create user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    user,
	})
}

// GetUserByID
// @Summary Get User by ID
// @Description Get a user by ID.
// @Param ID path string true "Get user by ID"
// @Produce application/json
// @Tags Users
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Router /users/{ID} [get]
func (ur *userController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "error",
			Message: "Invalid user ID",
		})
		return
	}
	user, err := ur.userRepository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "error",
			Message: `Unable to get user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "User found",
		Data:    user,
	})
}

// GetAllUser
// @Summary Get all Users
// @Description Get all Users.
// @Produce application/json
// @Tags Users
// @Success 200 {object} []models.Response{}
// @Failure 400 {object} []models.Response{}
// @Router /users [get]
func (ur *userController) GetAll(c *gin.Context) {
	users, err := ur.userRepository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "error",
			Message: `Unable to get users... ` + err.Error(),
		})
		return
	}
	if len(*users) == 0 {
		c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "No users found",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Users found",
		Data:    users,
	})
}
