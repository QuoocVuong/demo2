package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"product-catalog-service/models"
	"product-catalog-service/repository"
)

var nhomHangRepository repository.NhomHangRepository

func InitNhomHangHandler(repo repository.NhomHangRepository) {
	nhomHangRepository = repo
}

func GetAllNhomHangs(c *gin.Context) {
	nhomHangs, err := nhomHangRepository.GetAllNhomHang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nhomHangs)
}

func GetNhomHangByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	nhomHang, err := nhomHangRepository.GetNhomHangByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nhóm hàng không tồn tại"})
		return
	}
	c.JSON(http.StatusOK, nhomHang)
}

func CreateNhomHang(c *gin.Context) {
	var nhomHang models.NhomHang
	if err := c.ShouldBindJSON(&nhomHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := nhomHangRepository.CreateNhomHang(&nhomHang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nhomHang)
}

func UpdateNhomHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var nhomHang models.NhomHang
	if err := c.ShouldBindJSON(&nhomHang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := nhomHangRepository.UpdateNhomHang(uint(id), &nhomHang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nhomHang)
}

func DeleteNhomHang(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := nhomHangRepository.DeleteNhomHang(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xoá thành công"})
}
