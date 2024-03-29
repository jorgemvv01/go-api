package tests_controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/repositories"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.User{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.User{}); err != nil {
			t.Error(err)
		}
	}()

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	requestBody := `{"surname":"Jorge","lastname":"Villarreal"}`
	request := httptest.NewRequest("POST", "/users/create", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.POST("/users/create", userController.Create)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var user *models.User
	db.Last(&user)
	if user.Surname != "Jorge" || user.Lastname != "Villarreal" {
		t.Errorf("Unexpected user data: %v", user)
	}
}

func TestGetUserByID(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.User{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.User{}); err != nil {
			t.Error(err)
		}
	}()

	user := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}

	db.Create(&user)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	request := httptest.NewRequest("GET", "/users/1", nil)
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.GET("/users/:id", userController.GetByID)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	if err = json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Data == nil {
		t.Error(responseBody.Message)
	}

	data, ok := responseBody.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Bad data response structure")
	}
	if data["surname"] != "John" {
		t.Errorf("Surname does not match")
	}
	if data["lastname"] != "Doe" {
		t.Errorf("Lastname does not match")
	}
}

func TestGetAllUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.User{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.User{}); err != nil {
			t.Error(err)
		}
	}()

	user1 := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}
	user2 := models.User{
		Surname:  "George",
		Lastname: "Smith",
	}
	db.Create(&user1)
	db.Create(&user2)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	request := httptest.NewRequest("GET", "/users", nil)

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.GET("/users", userController.GetAll)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	if err = json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Data == nil {
		t.Error(responseBody.Message)
	}

	data, ok := responseBody.Data.([]interface{})
	if !ok {
		t.Errorf("Bad data response structure")
	}
	if len(data) != 2 {
		t.Errorf("No 2 users found")
	}
}

func TestUpdateUsers(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.User{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.User{}); err != nil {
			t.Error(err)
		}
	}()

	user := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}
	db.Create(&user)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	requestBody := `{"surname":"Jorge","lastname":"Villarreal"}`
	request := httptest.NewRequest("PUT", "/users/update/1", strings.NewReader(requestBody))

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.PUT("/users/update/:id", userController.Update)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	if err = json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Data == nil {
		t.Error(responseBody.Message)
	}

	data, ok := responseBody.Data.(map[string]interface{})

	if !ok {
		t.Errorf("Bad data response structure")
	}

	if data["surname"] != "Jorge" {
		t.Errorf("Surname does not match")
	}
	if data["lastname"] != "Villarreal" {
		t.Errorf("Lastname does not match")
	}
}

func TestDeleteUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.User{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.User{}); err != nil {
			t.Error(err)
		}
	}()

	user := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}
	db.Create(&user)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	request := httptest.NewRequest("DELETE", "/users/delete/1", nil)

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.DELETE("/users/delete/:id", userController.Delete)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	if err = json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	if responseBody.Message != "User deleted successfully" {
		t.Error("User could not be deleted successfully")
	}

}
