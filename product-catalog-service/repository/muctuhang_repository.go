package repository

import (
	"product-catalog-service/models"

	"gorm.io/gorm"
)

// Interface cho MucTuHangRepository
type MucTuHangRepository interface {
	GetAllMucTuHang() ([]models.MucTuHang, error)
	GetMucTuHangByID(uint) (*models.MucTuHang, error)
	CreateMucTuHang(*models.MucTuHang) error
	UpdateMucTuHang(uint, *models.MucTuHang) error
	DeleteMucTuHang(uint) error
}

// Triển khai MucTuHangRepository sử dụng GORM
type gormMucTuHangRepository struct {
	db *gorm.DB
}

// Hàm tạo gormMucTuHangRepository
func NewGormMucTuHangRepository(db *gorm.DB) MucTuHangRepository {
	return &gormMucTuHangRepository{db: db}
}

// Lấy tất cả mục từ hàng
func (r *gormMucTuHangRepository) GetAllMucTuHang() ([]models.MucTuHang, error) {
	var mucTuHangs []models.MucTuHang
	result := r.db.Find(&mucTuHangs)
	return mucTuHangs, result.Error
}

// Lấy mục từ hàng theo ID
func (r *gormMucTuHangRepository) GetMucTuHangByID(id uint) (*models.MucTuHang, error) {
	var mucTuHang models.MucTuHang
	result := r.db.First(&mucTuHang, id)
	return &mucTuHang, result.Error
}

// Tạo mục từ hàng mới
func (r *gormMucTuHangRepository) CreateMucTuHang(mucTuHang *models.MucTuHang) error {
	result := r.db.Create(mucTuHang)
	return result.Error
}

// Cập nhật thông tin mục từ hàng
func (r *gormMucTuHangRepository) UpdateMucTuHang(id uint, mucTuHang *models.MucTuHang) error {
	result := r.db.Model(&models.MucTuHang{}).Where("id = ?", id).Updates(mucTuHang)
	return result.Error
}

// Xoá mục từ hàng
func (r *gormMucTuHangRepository) DeleteMucTuHang(id uint) error {
	result := r.db.Delete(&models.MucTuHang{}, id)
	return result.Error
}
