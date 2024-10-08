package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"product-catalog-service/models"
	"product-catalog-service/repository"
)

var mucTuHangRepository repository.MucTuHangRepository

func InitMucTuHangHandler(repo repository.MucTuHangRepository) {
	mucTuHangRepository = repo
}

func GetAllMucTuHangs(c *gin.Context) {
	mucTuHangs, err := mucTuHangRepository.GetAllMucTuHang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mucTuHangs)
}

func GetMucTuHangByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mucTuHang, err := mucTuHangRepository.GetMucTuHangByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mục từ hàng không tồn tại"})
		return
	}
	c.JSON(http.StatusOK, mucTuHang)
}

func CreateMucTuHang(c *gin.Context) {
	var mucTuHang models.MucTuHang
	if err := c.ShouldBindJSON(&mucTuHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := mucTuHangRepository.CreateMucTuHang(&mucTuHang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mucTuHang)
}

func UpdateMucTuHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var mucTuHang models.MucTuHang
	if err := c.ShouldBindJSON(&mucTuHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := mucTuHangRepository.UpdateMucTuHang(uint(id), &mucTuHang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mucTuHang)
}

func DeleteMucTuHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := mucTuHangRepository.DeleteMucTuHang(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xoá thành công"})
}
