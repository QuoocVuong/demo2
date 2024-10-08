package repository

import (
	"gorm.io/gorm"

	"inventory-service/models"
)

// ... (KhoHangRepository interface and implementation)

type TonKhoRepository interface {
	GetAllTonKho() ([]models.TonKho, error)
	GetTonKhoByID(id uint) (*models.TonKho, error)
	CreateTonKho(tonKho *models.TonKho) error
	UpdateTonKho(tonKho *models.TonKho) error
	DeleteTonKho(id uint) error
}

type tonKhoRepository struct {
	db *gorm.DB
}

func NewTonKhoRepository(db *gorm.DB) *tonKhoRepository {
	return &tonKhoRepository{db: db}
}

func (r *tonKhoRepository) GetAllTonKho() ([]models.TonKho, error) {
	var tonKhos []models.TonKho
	err := r.db.Find(&tonKhos).Error
	return tonKhos, err
}

func (r *tonKhoRepository) GetTonKhoByID(id uint) (*models.TonKho, error) {
	var tonKho models.TonKho
	err := r.db.First(&tonKho, id).Error
	if err != nil {
		return nil, err
	}
	return &tonKho, nil
}

func (r *tonKhoRepository) CreateTonKho(tonKho *models.TonKho) error {
	return r.db.Create(tonKho).Error
}

func (r *tonKhoRepository) UpdateTonKho(tonKho *models.TonKho) error {
	return r.db.Save(tonKho).Error
}

func (r *tonKhoRepository) DeleteTonKho(id uint) error {
	return r.db.Delete(&models.TonKho{}, id).Error
}
