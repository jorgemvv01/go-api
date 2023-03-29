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

func TestCreateType(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}); err != nil {
			t.Fatal(err)
		}
	}()

	typeRepository := repositories.NewTypeRepository(db)
	typeController := controllers.NewTypeController(typeRepository)

	requestBody := `{"name":"New releases"}`
	request := httptest.NewRequest("POST", "/types/create", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.POST("/types/create", typeController.Create)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var movieType *models.Type
	db.Last(&movieType)
	if movieType.Name != "New releases" {
		t.Errorf("Unexpected type data: %v", movieType)
	}
}

func TestGetTypeByID(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{})
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}
	db.Create(&movieType)

	typeRepository := repositories.NewTypeRepository(db)
	typeController := controllers.NewTypeController(typeRepository)

	request := httptest.NewRequest("GET", "/types/1", nil)
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.GET("/types/:id", typeController.GetByID)
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
	if data["Name"] != "New releases" {
		t.Errorf("Name does not match")
	}
}

func TestGetAllType(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}); err != nil {
			t.Error(err)
		}
	}()

	type1 := models.Type{
		Name: "New releases",
	}
	type2 := models.Type{
		Name: "Regular movies",
	}
	db.Create(&type1)
	db.Create(&type2)

	typeRepository := repositories.NewTypeRepository(db)
	typeController := controllers.NewTypeController(typeRepository)

	request := httptest.NewRequest("GET", "/types", nil)
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.GET("/types", typeController.GetAll)
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
		t.Errorf("No 2 movie types found")
	}
}

func TestUpdateType(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}

	db.Create(&movieType)

	typeRepository := repositories.NewTypeRepository(db)
	typeController := controllers.NewTypeController(typeRepository)

	requestBody := `{"name":"Regular movies"}`
	request := httptest.NewRequest("PATCH", "/types/update/1", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.PATCH("/types/update/:id", typeController.Update)
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

	if data["Name"] != "Regular movies" {
		t.Errorf("Type name does not match")
	}
}

func TestDeleteType(t *testing.T) {
	router := gin.Default()
	db, err := setupDB(models.Type{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = dropTable(db, models.Type{}); err != nil {
			t.Error(err)
		}
	}()

	movieType := models.Type{
		Name: "New releases",
	}
	db.Create(&movieType)

	typeRepository := repositories.NewTypeRepository(db)
	typeController := controllers.NewTypeController(typeRepository)

	request := httptest.NewRequest("DELETE", "/types/delete/1", nil)
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.DELETE("/types/delete/:id", typeController.Delete)
	router.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var responseBody models.Response
	json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseBody.Message != "Type deleted successfully" {
		t.Error("Movie type could not be deleted successfully")
	}

}
