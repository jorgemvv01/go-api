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

func TestCreateMovie(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{}, models.Genre{}, models.Movie{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}, models.Genre{}, models.Movie{}); err != nil {
			t.Error(err)
		}
	}()

	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)

	movieType := models.Type{
		Name: "New releases",
	}
	genre := models.Genre{
		Name: "Action",
	}
	db.Create(&movieType)
	db.Create(&genre)

	requestBody := `{
	  "name": "Shazam! Fury of the Gods",
	  "overview": "Billy Batson and his foster siblings, who transform into superheroes by saying 'Shazam!', are forced to get back into action and fight the Daughters of Atlas, who they must stop from using a weapon that could destroy the world.",
	  "price": 11.25,
	  "type_id": 1,
      "genre_id": 1,
      "release_date": "2023-03-16"
	}`
	request := httptest.NewRequest("POST", "/movies/create", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.POST("/movies/create", movieController.Create)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
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
	if data["name"] != "Shazam! Fury of the Gods" {
		t.Errorf("Name does not match")
	}
	if data["price"] != 11.25 {
		t.Errorf("Price does not match")
	}
	if data["release_date"] != "2023-03-16" {
		t.Errorf("Release date does not match")
	}
}

func TestGetMovieByID(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{}, models.Genre{}, models.Movie{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}, models.Genre{}, models.Movie{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}
	genre := models.Genre{
		Name: "Science Fiction",
	}
	movie := models.Movie{
		Name:        "Avatar: The Way of Water",
		Overview:    "Set more than a decade after the events of the first film, learn the story of the Sully family (Jake, Neytiri, and their kids), the trouble that follows them, the lengths they go to keep each other safe, the battles they fight to stay alive, and the tragedies they endure.",
		Price:       10.25,
		TypeID:      1,
		GenreID:     1,
		ReleaseDate: "2022-12-15",
	}
	db.Create(&movieType)
	db.Create(&genre)
	db.Create(&movie)

	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)

	request := httptest.NewRequest("GET", "/movies/1", nil)
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.GET("/movies/:id", movieController.GetByID)
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
	if data["name"] != "Avatar: The Way of Water" {
		t.Errorf("Name does not match")
	}
	if data["price"] != 10.25 {
		t.Errorf("Price does not match")
	}
	if data["release_date"] != "2022-12-15" {
		t.Errorf("Release date does not match")
	}
}

func TestGetAllMovies(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{}, models.Genre{}, models.Movie{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}, models.Genre{}, models.Movie{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}
	genre1 := models.Genre{
		Name: "Science Fiction",
	}
	genre2 := models.Genre{
		Name: "Action",
	}
	movie1 := models.Movie{
		Name:        "Avatar: The Way of Water",
		Overview:    "Set more than a decade after the events of the first film, learn the story of the Sully family (Jake, Neytiri, and their kids), the trouble that follows them, the lengths they go to keep each other safe, the battles they fight to stay alive, and the tragedies they endure.",
		Price:       10.25,
		TypeID:      1,
		GenreID:     1,
		ReleaseDate: "2022-12-15",
	}
	movie2 := models.Movie{
		Name:        "Shazam! Fury of the Gods",
		Overview:    "Set more than a decade after the events of the first film, learn the story of the Sully family (Jake, Neytiri, and their kids), the trouble that follows them, the lengths they go to keep each other safe, the battles they fight to stay alive, and the tragedies they endure.",
		Price:       11.25,
		TypeID:      1,
		GenreID:     2,
		ReleaseDate: "2023-03-16",
	}
	db.Create(&movieType)
	db.Create(&genre1)
	db.Create(&genre2)
	db.Create(&movie1)
	db.Create(&movie2)

	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)

	request := httptest.NewRequest("GET", "/movies", nil)

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.GET("/movies", movieController.GetAll)
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
		t.Errorf("No 2 movies found")
	}
}

func TestUpdateMovie(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{}, models.Genre{}, models.Movie{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}, models.Genre{}, models.Movie{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}
	genre1 := models.Genre{
		Name: "Science Fiction",
	}
	genre2 := models.Genre{
		Name: "Action",
	}
	movie := models.Movie{
		Name:        "Avatar: The Way of Water",
		Overview:    "Set more than a decade after the events of the first film, learn the story of the Sully family (Jake, Neytiri, and their kids), the trouble that follows them, the lengths they go to keep each other safe, the battles they fight to stay alive, and the tragedies they endure.",
		Price:       10.25,
		TypeID:      1,
		GenreID:     1,
		ReleaseDate: "2022-12-15",
	}
	db.Create(&movieType)
	db.Create(&genre1)
	db.Create(&genre2)
	db.Create(&movie)

	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)

	requestBody := `{
	  "name": "Shazam! Fury of the Gods",
	  "overview": "Billy Batson and his foster siblings, who transform into superheroes by saying 'Shazam!', are forced to get back into action and fight the Daughters of Atlas, who they must stop from using a weapon that could destroy the world.",
	  "price": 11.25,
	  "type_id": 1,
      "genre_id": 2,
      "release_date": "2023-03-16"
	}`
	request := httptest.NewRequest("PUT", "/movies/update/1", strings.NewReader(requestBody))

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.PUT("/movies/update/:id", movieController.Update)
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

	if data["name"] != "Shazam! Fury of the Gods" {
		t.Errorf("Name does not match")
	}
	if data["price"] != 11.25 {
		t.Errorf("Price does not match")
	}
	if data["release_date"] != "2023-03-16" {
		t.Errorf("Release date does not match")
	}
}

func TestDeleteMovie(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{}, models.Genre{}, models.Movie{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}, models.Genre{}, models.Movie{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}
	genre := models.Genre{
		Name: "Science Fiction",
	}
	movie := models.Movie{
		Name:        "Avatar: The Way of Water",
		Overview:    "Set more than a decade after the events of the first film, learn the story of the Sully family (Jake, Neytiri, and their kids), the trouble that follows them, the lengths they go to keep each other safe, the battles they fight to stay alive, and the tragedies they endure.",
		Price:       10.25,
		TypeID:      1,
		GenreID:     1,
		ReleaseDate: "2022-12-15",
	}
	db.Create(&movieType)
	db.Create(&genre)
	db.Create(&movie)

	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)

	request := httptest.NewRequest("DELETE", "/movies/delete/1", nil)

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.DELETE("/movies/delete/:id", movieController.Delete)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseBody.Message != "Movie deleted successfully" {
		t.Error("Movie could not be deleted successfully")
	}

}
