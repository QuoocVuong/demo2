package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NhomHang struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TenNhom string `json:"ten_nhom"`
}

type SanPham struct {
	ID                           uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	MaHang                       string   `json:"ma_hang"`
	TenMuc                       string   `json:"ten_muc"`
	NhomHangID                   uint     `json:"nhom_hang_id"`
	DonViDoDinh                  string   `json:"don_vi_do_dinh"`
	VoHieuHoa                    bool     `json:"vo_hieu_hoa"`
	ChoPhepKhoanThayThe          bool     `json:"cho_phep_khoan_thay_the"`
	DuyTriHangTonKho             bool     `json:"duy_tri_hang_ton_kho"`
	BaoQumCacMatHangTrongSanXuat bool     `json:"bao_qum_cac_mat_hang_trong_san_xuat"`
	CoPhieuMoDau                 string   `json:"co_phieu_mo_dau"`
	DinhGia                      float64  `json:"dinh_gia"`
	TyGiaBanHangTauChuan         float64  `json:"ty_gia_ban_hang_tau_chuan"`
	LaChiDinhTaiSan              bool     `json:"la_chi_dinh_tai_san"`
	NhomHang                     NhomHang `gorm:"foreignKey:NhomHangID" json:"nhom_hang"` // Add this for eager loading
}

type MucTuHang struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	SanPhamID uint    `json:"san_pham_id"`
	TenMucTu  string  `json:"ten_muc_tu"`
	SanPham   SanPham `gorm:"foreignKey:SanPhamID" json:"san_pham"`
}

type KhoHang struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TenKho string `json:"ten_kho"`
}

type TonKho struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SanPhamID   uint      `json:"san_pham_id"`
	KhoHangID   uint      `json:"kho_hang_id"`
	SoLuong     int       `json:"so_luong"`
	NgayCapNhat time.Time `json:"ngay_cap_nhat"`
	SanPham     SanPham   `gorm:"foreignKey:SanPhamID" json:"san_pham"`
	KhoHang     KhoHang   `gorm:"foreignKey:KhoHangID" json:"kho_hang"`
}

var db *gorm.DB

func main() {
	// Database connection
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo2?charset=utf8mb4&parseTime=True&loc=Local" // Replace with your DB credentials
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&NhomHang{}, &SanPham{}, &MucTuHang{}, &KhoHang{}, &TonKho{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

	r := gin.Default()

	// Nhóm hàng routes
	nhomHangRoutes(r)

	// Sản phẩm routes
	sanPhamRoutes(r)
	// ... (other routes for MucTuHang, KhoHang, TonKho)
	mucTuHangRoutes(r) // Add MucTuHang routes
	tonKhoRoutes(r)    // Add TonKho routes
	khoHangRoutes(r)   // Add KhoHang routes

	r.Run(":8080")

}

// ... (CRUD functions for NhomHang, SanPham, MucTuHang, KhoHang, TonKho as in the previous examples. Add Preload to SanPham GET requests.)

func nhomHangRoutes(r *gin.Engine) {
	r.GET("/nhom-hang", getNhomHangs)
	r.GET("/nhom-hang/:id", getNhomHang)
	r.POST("/nhom-hang", createNhomHang)
	r.PUT("/nhom-hang/:id", updateNhomHang)
	r.DELETE("/nhom-hang/:id", deleteNhomHang)
}

func sanPhamRoutes(r *gin.Engine) {
	r.GET("/san-pham", getSanPhams)
	r.GET("/san-pham/:id", getSanPham)
	r.POST("/san-pham", createSanPham)
	r.PUT("/san-pham/:id", updateSanPham)
	r.DELETE("/san-pham/:id", deleteSanPham)
}
func mucTuHangRoutes(r *gin.Engine) {
	r.GET("/muc-tu-hang", getMucTuHangs)
	r.GET("/muc-tu-hang/:id", getMucTuHang)
	r.POST("/muc-tu-hang", createMucTuHang)
	r.PUT("/muc-tu-hang/:id", updateMucTuHang)
	r.DELETE("/muc-tu-hang/:id", deleteMucTuHang)
}

func tonKhoRoutes(r *gin.Engine) {
	r.GET("/ton-kho", getTonKhos)
	r.GET("/ton-kho/:id", getTonKho)
	r.POST("/ton-kho", createTonKho)
	r.PUT("/ton-kho/:id", updateTonKho)
	r.DELETE("/ton-kho/:id", deleteTonKho)
}

func khoHangRoutes(r *gin.Engine) {
	r.GET("/kho-hang", getKhoHangs)
	r.GET("/kho-hang/:id", getKhoHang)
	r.POST("/kho-hang", createKhoHang)
	r.PUT("/kho-hang/:id", updateKhoHang)
	r.DELETE("/kho-hang/:id", deleteKhoHang)
}

// ... (rest of the CRUD functions)

func getSanPhams(c *gin.Context) {
	var sanPhams []SanPham
	db.Preload("NhomHang").Find(&sanPhams) // Eager load NhomHang
	c.JSON(http.StatusOK, sanPhams)
}

// Example using Preload for SanPham GET request
func getSanPham(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var sanPham SanPham
	if err := db.Preload("NhomHang").First(&sanPham, id).Error; err != nil { // Preload NhomHang
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, sanPham)
}

// Example CRUD function for NhomHang (others will be similar)
func getNhomHangs(c *gin.Context) {
	var nhomHangs []NhomHang
	db.Find(&nhomHangs)
	c.JSON(http.StatusOK, nhomHangs)
}

func getNhomHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var nhomHang NhomHang
	if err := db.First(&nhomHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, nhomHang)
}

func createNhomHang(c *gin.Context) {
	var nhomHang NhomHang
	if err := c.ShouldBindJSON(&nhomHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&nhomHang)
	c.JSON(http.StatusCreated, nhomHang)
}

func updateNhomHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var nhomHang NhomHang
	if err := db.First(&nhomHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	if err := c.ShouldBindJSON(&nhomHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&nhomHang)
	c.JSON(http.StatusOK, nhomHang)

}

func deleteNhomHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var nhomHang NhomHang
	if err := db.First(&nhomHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	fmt.Println("Deleting nhomHang:", nhomHang) // Add debug print

	db.Delete(&nhomHang)

	c.JSON(http.StatusOK, gin.H{"message": "record deleted successfully"})
}

// ... (other imports and struct definitions)

func createSanPham(c *gin.Context) {
	var sanPham SanPham
	if err := c.ShouldBindJSON(&sanPham); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if NhomHang exists (important for foreign key constraint)
	var nhomHang NhomHang
	if err := db.First(&nhomHang, sanPham.NhomHangID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nhom_hang_id"})
		return
	}

	if err := db.Create(&sanPham).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, sanPham)
}

// ... other imports and struct definitions

func updateSanPham(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sanPham SanPham
	if err := db.First(&sanPham, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&sanPham); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: Check if NhomHang exists if you're allowing updates to nhom_hang_id
	if sanPham.NhomHangID != 0 { // Check if nhom_hang_id is being updated
		var nhomHang NhomHang
		if err := db.First(&nhomHang, sanPham.NhomHangID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nhom_hang_id"})
			return
		}
	}

	db.Save(&sanPham) // This will update the record
	c.JSON(http.StatusOK, sanPham)
}

func deleteSanPham(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sanPham SanPham
	if err := db.First(&sanPham, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&sanPham)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
func getMucTuHangs(c *gin.Context) {
	var mucTuHangs []MucTuHang
	db.Preload("SanPham").Find(&mucTuHangs)
	c.JSON(http.StatusOK, mucTuHangs)
}

func getMucTuHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var mucTuHang MucTuHang
	if err := db.Preload("SanPham").First(&mucTuHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, mucTuHang)
}

func createMucTuHang(c *gin.Context) {
	var mucTuHang MucTuHang
	if err := c.ShouldBindJSON(&mucTuHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if SanPham exists
	var sanPham SanPham
	if err := db.First(&sanPham, mucTuHang.SanPhamID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid san_pham_id"})
		return
	}

	db.Create(&mucTuHang)
	c.JSON(http.StatusCreated, mucTuHang)
}

// ... (other imports and struct definitions)

// MucTuHang CRUD functions

func updateMucTuHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var mucTuHang MucTuHang
	if err := db.First(&mucTuHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&mucTuHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: Check if SanPham exists if you allow san_pham_id updates
	if mucTuHang.SanPhamID != 0 {
		var sanPham SanPham
		if err := db.First(&sanPham, mucTuHang.SanPhamID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid san_pham_id"})
			return
		}
	}

	db.Save(&mucTuHang)
	c.JSON(http.StatusOK, mucTuHang)
}

func deleteMucTuHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var mucTuHang MucTuHang
	if err := db.First(&mucTuHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&mucTuHang)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}

// TonKho CRUD functions

func getTonKhos(c *gin.Context) {
	var tonKhos []TonKho
	db.Preload("SanPham").Preload("KhoHang").Find(&tonKhos) // Preload both SanPham and KhoHang
	c.JSON(http.StatusOK, tonKhos)
}

func getTonKho(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var tonKho TonKho
	if err := db.Preload("SanPham").Preload("KhoHang").First(&tonKho, id).Error; err != nil { // Preload for efficient querying
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, tonKho)
}

func createTonKho(c *gin.Context) {
	var tonKho TonKho
	if err := c.ShouldBindJSON(&tonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if both SanPham and KhoHang exist
	var sanPham SanPham
	if err := db.First(&sanPham, tonKho.SanPhamID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid san_pham_id"})
		return
	}

	var khoHang KhoHang
	if err := db.First(&khoHang, tonKho.KhoHangID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kho_hang_id"})
		return
	}

	db.Create(&tonKho)
	c.JSON(http.StatusCreated, tonKho)
}

func updateTonKho(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var tonKho TonKho
	if err := db.First(&tonKho, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&tonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if SanPham and KhoHang exist (if these fields are being updated)
	if tonKho.SanPhamID != 0 {
		var sanPham SanPham
		if err := db.First(&sanPham, tonKho.SanPhamID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid san_pham_id"})
			return
		}
	}

	if tonKho.KhoHangID != 0 {
		var khoHang KhoHang
		if err := db.First(&khoHang, tonKho.KhoHangID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kho_hang_id"})
			return
		}
	}
	fmt.Println("TonKho before save:", tonKho) // Debug print statement

	db.Save(&tonKho)
	fmt.Println("TonKho after save:", tonKho) // Debug print statement
	c.JSON(http.StatusOK, tonKho)

}

func deleteTonKho(c *gin.Context) {
	// ... (implementation similar to deleteSanPham and deleteMucTuHang)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var tonKho TonKho
	if err := db.First(&tonKho, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&tonKho)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}

// ... (other imports and struct definitions)

// KhoHang CRUD functions

func getKhoHangs(c *gin.Context) {
	var khoHangs []KhoHang
	db.Find(&khoHangs)
	c.JSON(http.StatusOK, khoHangs)
}

func getKhoHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var khoHang KhoHang
	if err := db.First(&khoHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, khoHang)
}

func createKhoHang(c *gin.Context) {
	var khoHang KhoHang
	if err := c.ShouldBindJSON(&khoHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&khoHang)
	c.JSON(http.StatusCreated, khoHang)
}

func updateKhoHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var khoHang KhoHang
	if err := db.First(&khoHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&khoHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&khoHang)
	c.JSON(http.StatusOK, khoHang)
}

func deleteKhoHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var khoHang KhoHang
	if err := db.First(&khoHang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&khoHang)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}

// ... (KhoHang CRUD functions - implement these similarly)
// ... (CRUD functions for other entities)
