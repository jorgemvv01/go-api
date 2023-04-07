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

func TestCreateGenre(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Genre{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Genre{}); err != nil {
			t.Fatal(err)
		}
	}()

	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)

	requestBody := `{"name":"Action"}`
	request := httptest.NewRequest("POST", "/genres/create", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.POST("/genres/create", genreController.Create)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var genre *models.Genre
	db.Last(&genre)
	if genre.Name != "Action" {
		t.Errorf("Unexpected genre name: %v", genre.Name)
	}
}

func TestGetGenreByID(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Genre{})
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err = dropTable(db, models.Genre{}); err != nil {
			t.Error(err)
		}
	}()

	genre := models.Genre{
		Name: "Action",
	}
	db.Create(&genre)

	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)

	request := httptest.NewRequest("GET", "/genres/1", nil)
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.GET("/genres/:id", genreController.GetByID)
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
	if data["name"] != "Action" {
		t.Errorf("Name does not match")
	}
}

func TestGetAllGenres(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Genre{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Genre{}); err != nil {
			t.Error(err)
		}
	}()

	genre1 := models.Genre{
		Name: "Action",
	}
	genre2 := models.Genre{
		Name: "Comedy",
	}
	db.Create(&genre1)
	db.Create(&genre2)

	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)

	request := httptest.NewRequest("GET", "/genres", nil)
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.GET("/genres", genreController.GetAll)
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
		t.Errorf("No 2 movie genres found")
	}
}

func TestUpdateGenre(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Genre{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Genre{}); err != nil {
			t.Error(err)
		}
	}()

	genre := models.Genre{
		Name: "Action",
	}

	db.Create(&genre)

	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)

	requestBody := `{"name":"Comedy"}`
	request := httptest.NewRequest("PUT", "/genres/update/1", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.PUT("/genres/update/:id", genreController.Update)
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

	if data["name"] != "Comedy" {
		t.Errorf("Genre name does not match")
	}
}

func TestDeleteGenre(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Genre{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Genre{}); err != nil {
			t.Error(err)
		}
	}()

	genre := models.Genre{
		Name: "Action",
	}
	db.Create(&genre)

	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)

	request := httptest.NewRequest("DELETE", "/genres/delete/1", nil)
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.DELETE("/genres/delete/:id", genreController.Delete)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseBody.Message != "Genre deleted successfully" {
		t.Error("Movie genre could not be deleted successfully")
	}

}
