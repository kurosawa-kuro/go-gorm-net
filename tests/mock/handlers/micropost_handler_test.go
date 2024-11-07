package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"go-gorm-net/internal/handlers"
	"go-gorm-net/internal/models"
	"go-gorm-net/internal/services/mock_services" // Import the generated mock package

	"github.com/golang/mock/gomock"
)

func TestHandleMicroposts_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockMicropostServiceInterface(ctrl)
	handler := handlers.NewMicropostHandler(mockService)

	// Set up expected behavior
	mockService.EXPECT().GetAll().Return([]models.Micropost{}, nil)

	req, err := http.NewRequest(http.MethodGet, "/microposts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.HandleMicroposts(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Unmarshal the response body
	var actual []models.Micropost
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expected := []models.Micropost{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestHandleMicroposts_Post(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockMicropostServiceInterface(ctrl)
	handler := handlers.NewMicropostHandler(mockService)

	micropost := models.Micropost{ID: 1, Title: "Test content"}
	mockService.EXPECT().Create(&micropost).Return(nil)

	body, _ := json.Marshal(micropost)
	req, err := http.NewRequest(http.MethodPost, "/microposts", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.HandleMicroposts(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Unmarshal the response body
	var actual models.Micropost
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expected := micropost
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
