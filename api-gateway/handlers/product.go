package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Q.Vuong/demo2/proto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductClient pb.ProductServiceClient
}

func (h *ProductHandler) GetSanPham(c *gin.Context) {
	// Lấy ID từ URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Gọi gRPC đến product service
	res, err := h.ProductClient.GetSanPham(context.Background(), &pb.GetSanPhamRequest{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, res.GetSanPham())
}
func (h *ProductHandler) GetAllSanPhams(c *gin.Context) {
	// Gọi gRPC đến product service
	res, err := h.ProductClient.GetAllSanPhams(context.Background(), &pb.GetAllSanPhamsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, res.GetSanPhams())
}
