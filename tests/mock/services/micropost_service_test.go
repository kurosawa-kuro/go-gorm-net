package services

import (
	"testing"

	"go-gorm-net/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMicropostService is a mock implementation of the MicropostService
type MockMicropostService struct {
	mock.Mock
}

func (m *MockMicropostService) GetAll() ([]models.Micropost, error) {
	args := m.Called()
	return args.Get(0).([]models.Micropost), args.Error(1)
}

func (m *MockMicropostService) Create(micropost *models.Micropost) error {
	args := m.Called(micropost)
	return args.Error(0)
}

func (m *MockMicropostService) GetByID(id int) (*models.Micropost, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Micropost), args.Error(1)
}

func TestMicropostService_GetAll(t *testing.T) {
	mockService := new(MockMicropostService)

	mockMicroposts := []models.Micropost{
		{ID: 1, Title: "Hello World"},
	}
	mockService.On("GetAll").Return(mockMicroposts, nil)

	microposts, err := mockService.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, mockMicroposts, microposts)
	mockService.AssertExpectations(t)
}
