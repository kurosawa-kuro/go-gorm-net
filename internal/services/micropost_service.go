package services

import (
	"go-gorm-net/internal/models"
	"go-gorm-net/pkg/database"
)

type MicropostService struct{}

func NewMicropostService() *MicropostService {
	return &MicropostService{}
}

func (s *MicropostService) GetAll() ([]models.Micropost, error) {
	var microposts []models.Micropost
	// idが降順になるように取得
	result := database.DB.Order("id DESC").Find(&microposts)
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
