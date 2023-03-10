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
	Update(c *gin.Context)
	Delete(c *gin.Context)
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
// @Failure 500 {object} models.Response{}
// @Router /users/create [post]
func (uc *userController) Create(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	if err := uc.userRepository.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to create user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
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
// @Failure 500 {object} models.Response{}
// @Router /users/{ID} [get]
func (ur *userController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: "Invalid user ID",
		})
		return
	}
	user, err := ur.userRepository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "User found",
		Data:    user,
	})
}

// GetAllUser
// @Summary Get all Users
// @Description Get all Users.
// @Produce application/json
// @Tags Users
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /users [get]
func (ur *userController) GetAll(c *gin.Context) {
	users, err := ur.userRepository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to get users... ` + err.Error(),
		})
		return
	}
	if len(*users) == 0 {
		c.JSON(http.StatusOK, models.Response{
			Status:  "Success",
			Message: "No users found",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "Users found",
		Data:    users,
	})
}

// UpdateUser
// @Summary Update User
// @Description Update User by ID.
// @Produce application/json
// @Param ID path string true "Update user by ID"
// @Param tags body models.UserRequest true "Create user"
// @Tags Users
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /users/update/{ID} [patch]
func (uc *userController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: "Invalid user ID",
		})
		return
	}
	var user *models.User
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: `Invalid request body... ` + err.Error(),
		})
		return
	}
	if user, err = uc.userRepository.Update(uint(id), user); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to update user... ` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "User updated successfully",
		Data:    user,
	})
}

// DeleteUser
// @Summary Delete User
// @Description Delete User by ID.
// @Produce application/json
// @Param ID path string true "Delete user by ID"
// @Tags Users
// @Success 200 {object} models.Response{}
// @Failure 400 {object} models.Response{}
// @Failure 500 {object} models.Response{}
// @Router /users/delete/{ID} [delete]
func (uc *userController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Invalid user ID",
		})
		return
	}
	if err = uc.userRepository.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "Error",
			Message: `Unable to delete user...` + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "Success",
		Message: "User deleted successfully",
	})
}
