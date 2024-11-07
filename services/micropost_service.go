package services

import (
	"go-gorm-net/database"
	"go-gorm-net/models"
)

type MicropostService struct{}

func NewMicropostService() *MicropostService {
	return &MicropostService{}
}

func (s *MicropostService) GetAll() ([]models.Micropost, error) {
	var microposts []models.Micropost
	result := database.DB.Find(&microposts)
	return microposts, result.Error
}

func (s *MicropostService) Create(micropost *models.Micropost) error {
	return database.DB.Create(micropost).Error
}

func (s *MicropostService) GetByID(id int) (*models.Micropost, error) {
	var micropost models.Micropost
	result := database.DB.First(&micropost, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &micropost, nil
}
