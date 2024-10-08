package handlers

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"models"
	"repository"
)

var khoHangRepo repository.KhoHangRepository

func GetKhoHangs(c *gin.Context) {
	khoHangs, err := khoHangRepo.GetAllKhoHang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve KhoHangs"})
		return
	}
	c.JSON(http.StatusOK, khoHangs)
}

func GetKhoHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	khoHang, err := khoHangRepo.GetKhoHangByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve KhoHang"})
		}
		return
	}

	c.JSON(http.StatusOK, khoHang)
}

//	func CreateKhoHang(c *gin.Context) {
//		var khoHang models.KhoHang
//		if err := c.ShouldBindJSON(&khoHang); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		err := khoHangRepo.CreateKhoHang(&khoHang)
//
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create KhoHang"})
//			return
//		}
//
//		c.JSON(http.StatusCreated, khoHang)
//	}
func CreateKhoHang(c *gin.Context) {
	var khoHang models.KhoHang
	if err := c.ShouldBindJSON(&khoHang); err != nil {
		// Kiểm tra lỗi ràng buộc dữ liệu
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra các trường bắt buộc (tùy theo logic)
	if khoHang.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	err := khoHangRepo.CreateKhoHang(&khoHang)

	if err != nil {
		// Xử lý các lỗi khác nhau từ repository
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create KhoHang"})
		return
	}

	c.JSON(http.StatusCreated, khoHang)
}
func UpdateKhoHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	khoHang, err := khoHangRepo.GetKhoHangByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve KhoHang"})
		}
		return
	}

	if err := c.ShouldBindJSON(&khoHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = khoHangRepo.UpdateKhoHang(khoHang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update KhoHang"})
		return
	}

	c.JSON(http.StatusOK, khoHang)
}

func DeleteKhoHang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = khoHangRepo.DeleteKhoHang(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete KhoHang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
