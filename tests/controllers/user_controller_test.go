package tests_controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github/jorgemvv01/go-api/controllers"
	"github/jorgemvv01/go-api/models"
	"github/jorgemvv01/go-api/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect memory database: %v", err)
	}
	if err = db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate User model: %v", err)
	}
	return db, nil
}

func dropUserTable(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		return fmt.Errorf("failed to drop User table: %v", err)
	}
	return nil
}

func TestCreateUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer func() {
		if err = dropUserTable(db); err != nil {
			t.Errorf("error cleaning User table: %v", err)
		}
	}()

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	requestBody := `{"surname":"Jorge","lastname":"Villarreal"}`
	request, err := http.NewRequest("POST", "/users/create", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
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
	db, err := setupDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer func() {
		if err = dropUserTable(db); err != nil {
			t.Errorf("error cleaning User table: %v", err)
		}
	}()

	user := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}

	db.Create(&user)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	request, err := http.NewRequest("GET", "/users/1", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.GET("/users/:id", userController.GetByID)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseBody.Data == nil {
		t.Error(responseBody.Message)
	}

	data, ok := responseBody.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Bad data response structure")
	}
	if data["Surname"] != "John" {
		t.Errorf("Surname does not match")
	}
	if data["Lastname"] != "Doe" {
		t.Errorf("Lastname does not match")
	}
}

func TestGetAllUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer func() {
		if err = dropUserTable(db); err != nil {
			t.Errorf("error cleaning User table: %v", err)
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

	request, err := http.NewRequest("GET", "/users", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.GET("/users", userController.GetAll)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
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

func TestUpdateUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer func() {
		if err = dropUserTable(db); err != nil {
			t.Errorf("error cleaning User table: %v", err)
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
	request, err := http.NewRequest("PATCH", "/users/update/1", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.PATCH("/users/update/:id", userController.Update)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseBody.Data == nil {
		t.Error(responseBody.Message)
	}

	data, ok := responseBody.Data.(map[string]interface{})

	if !ok {
		t.Errorf("Bad data response structure")
	}

	if data["Surname"] != "Jorge" {
		t.Errorf("Surname does not match")
	}
	if data["Lastname"] != "Villarreal" {
		t.Errorf("Lastname does not match")
	}
}

func TestDeleteUser(t *testing.T) {
	router := gin.Default()
	db, err := setupDB()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}
	defer func() {
		if err = dropUserTable(db); err != nil {
			t.Errorf("error cleaning User table: %v", err)
		}
	}()

	user := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}
	db.Create(&user)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	request, err := http.NewRequest("DELETE", "/users/delete/1", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.DELETE("/users/delete/:id", userController.Delete)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseBody.Message != "User deleted successfully" {
		t.Error("User could not be deleted successfully")
	}

}
