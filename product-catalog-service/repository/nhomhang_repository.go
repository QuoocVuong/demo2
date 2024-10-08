package repository

import (
	"product-catalog-service/models"

	"gorm.io/gorm"
)

// Interface cho NhomHangRepository
type NhomHangRepository interface {
	GetAllNhomHang() ([]models.NhomHang, error)
	GetNhomHangByID(uint) (*models.NhomHang, error)
	CreateNhomHang(*models.NhomHang) error
	UpdateNhomHang(uint, *models.NhomHang) error
	DeleteNhomHang(uint) error
}

// Triển khai NhomHangRepository sử dụng GORM
type gormNhomHangRepository struct {
	db *gorm.DB
}

// Hàm tạo gormNhomHangRepository
func NewGormNhomHangRepository(db *gorm.DB) NhomHangRepository {
	return &gormNhomHangRepository{db: db}
}

// Lấy tất cả nhóm hàng
func (r *gormNhomHangRepository) GetAllNhomHang() ([]models.NhomHang, error) {
	var nhomHangs []models.NhomHang
	result := r.db.Find(&nhomHangs)
	return nhomHangs, result.Error
}

// Lấy nhóm hàng theo ID
func (r *gormNhomHangRepository) GetNhomHangByID(id uint) (*models.NhomHang, error) {
	var nhomHang models.NhomHang
	result := r.db.First(&nhomHang, id)
	return &nhomHang, result.Error
}

// Tạo nhóm hàng mới
func (r *gormNhomHangRepository) CreateNhomHang(nhomHang *models.NhomHang) error {
	result := r.db.Create(nhomHang)
	return result.Error
}

// Cập nhật thông tin nhóm hàng
func (r *gormNhomHangRepository) UpdateNhomHang(id uint, nhomHang *models.NhomHang) error {
	result := r.db.Model(&models.NhomHang{}).Where("id = ?", id).Updates(nhomHang)
	return result.Error
}

// Xoá nhóm hàng
func (r *gormNhomHangRepository) DeleteNhomHang(id uint) error {
	result := r.db.Delete(&models.NhomHang{}, id)
	return result.Error
}
