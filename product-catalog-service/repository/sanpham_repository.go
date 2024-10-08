package repository

import (
	"product-catalog-service/models"

	"gorm.io/gorm"
)

// Interface cho SanPhamRepository
type SanPhamRepository interface {
	GetAllSanPham() ([]models.SanPham, error)
	GetSanPhamByID(uint) (*models.SanPham, error)
	CreateSanPham(*models.SanPham) error
	UpdateSanPham(uint, *models.SanPham) error
	DeleteSanPham(uint) error
}

// Triển khai SanPhamRepository sử dụng GORM
type gormSanPhamRepository struct {
	db *gorm.DB
}

// Hàm tạo gormSanPhamRepository
func NewGormSanPhamRepository(db *gorm.DB) SanPhamRepository {
	return &gormSanPhamRepository{db: db}
}

// Lấy tất cả sản phẩm
func (r *gormSanPhamRepository) GetAllSanPham() ([]models.SanPham, error) {
	var sanPhams []models.SanPham
	result := r.db.Find(&sanPhams)
	return sanPhams, result.Error
}

// Lấy sản phẩm theo ID
func (r *gormSanPhamRepository) GetSanPhamByID(id uint) (*models.SanPham, error) {
	var sanPham models.SanPham
	result := r.db.First(&sanPham, id)
	return &sanPham, result.Error
}

// Tạo sản phẩm mới
func (r *gormSanPhamRepository) CreateSanPham(sanPham *models.SanPham) error {
	result := r.db.Create(sanPham)
	return result.Error
}

// Cập nhật thông tin sản phẩm
func (r *gormSanPhamRepository) UpdateSanPham(id uint, sanPham *models.SanPham) error {
	result := r.db.Model(&models.SanPham{}).Where("id = ?", id).Updates(sanPham)
	return result.Error
}

// Xoá sản phẩm
func (r *gormSanPhamRepository) DeleteSanPham(id uint) error {
	result := r.db.Delete(&models.SanPham{}, id)
	return result.Error
}
