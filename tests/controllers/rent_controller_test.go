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

func TestCreateRent(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{}, models.Genre{}, models.Movie{}, models.User{}, models.Rent{}, models.MovieRent{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}, models.Genre{}, models.Movie{}, models.User{}, models.Rent{}, models.MovieRent{}); err != nil {
			t.Error(err)
		}
	}()

	movieType1 := models.Type{
		Name: "New releases",
	}
	movieType2 := models.Type{
		Name: "Regular movies",
	}
	movieType3 := models.Type{
		Name: "Old movies",
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
		Price:       11.25,
		TypeID:      1,
		GenreID:     1,
		ReleaseDate: "2022-12-15",
	}
	movie2 := models.Movie{
		Name:        "Rambo",
		Overview:    "When governments fail to act on behalf of captive missionaries, ex-Green Beret John James Rambo sets aside his peaceful existence along the Salween River in a war-torn region of Thailand to take action. Although he's still haunted by violent memories of his time as a U.S. soldier during the Vietnam War, Rambo can hardly turn his back on the aid workers who so desperately need his help.",
		Price:       9.78,
		TypeID:      3,
		GenreID:     2,
		ReleaseDate: "2008-01-25",
	}
	user := models.User{
		Surname:  "John",
		Lastname: "Doe",
	}

	db.Create(&movieType1)
	db.Create(&movieType2)
	db.Create(&movieType3)
	db.Create(&genre1)
	db.Create(&genre2)
	db.Create(&movie1)
	db.Create(&movie2)
	db.Create(&user)

	rentRepository := repositories.NewRentRepository(db)
	rentController := controllers.NewRentController(rentRepository)

	requestBody := `{
      "user_id": 1,
	  "movie_ids": [
		1,2
	  ],
	  "start_date": "2023-04-07",
      "end_date": "2023-04-15"
	}`
	request := httptest.NewRequest("POST", "/rent/create", strings.NewReader(requestBody))

	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.POST("/rent/create", rentController.Create)
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
	if data["user_id"] != 1.0 {
		t.Errorf("User does not match")
	}
	if data["total"] != 171.174 {
		t.Errorf("Total does not match")
	}
	if data["start_date"] != "2023-04-07" {
		t.Errorf("Start date does not match")
	}
	if data["end_date"] != "2023-04-15" {
		t.Errorf("End date does not match")
	}
}
