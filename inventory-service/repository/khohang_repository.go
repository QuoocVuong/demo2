package repository

import (
	"inventory-service/models"

	"gorm.io/gorm"
)

type KhoHangRepository interface {
	GetAllKhoHang() ([]models.KhoHang, error)
	GetKhoHangByID(id uint) (*models.KhoHang, error)
	CreateKhoHang(khoHang *models.KhoHang) error
	UpdateKhoHang(khoHang *models.KhoHang) error
	DeleteKhoHang(id uint) error
}

type khoHangRepository struct {
	db *gorm.DB
}

func NewKhoHangRepository(db *gorm.DB) *khoHangRepository {
	return &khoHangRepository{db: db}
}

func (r *khoHangRepository) GetAllKhoHang() ([]models.KhoHang, error) {
	var khoHangs []models.KhoHang
	err := r.db.Find(&khoHangs).Error
	return khoHangs, err
}

func (r *khoHangRepository) GetKhoHangByID(id uint) (*models.KhoHang, error) {
	var khoHang models.KhoHang
	err := r.db.First(&khoHang, id).Error
	if err != nil {
		return nil, err
	}
	return &khoHang, nil
}

func (r *khoHangRepository) CreateKhoHang(khoHang *models.KhoHang) error {
	return r.db.Create(khoHang).Error
}

func (r *khoHangRepository) UpdateKhoHang(khoHang *models.KhoHang) error {
	return r.db.Save(khoHang).Error
}

func (r *khoHangRepository) DeleteKhoHang(id uint) error {
	return r.db.Delete(models.KhoHang{}, id).Error
}
