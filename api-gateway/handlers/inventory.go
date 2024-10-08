package handlers

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Q.Vuong/demo2/proto" // Import package proto của inventory service
	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	InventoryClient pb.InventoryServiceClient
}

func (h *InventoryHandler) GetAllTonKhos(c *gin.Context) {
	// Gọi gRPC đến inventory service
	res, err := h.InventoryClient.GetAllTonKhos(context.Background(), &pb.GetAllTonKhosRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, res.GetTonKhos())
}

func (h *InventoryHandler) GetTonKho(c *gin.Context) {
	// Lấy ID từ URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Gọi gRPC đến inventory service
	res, err := h.InventoryClient.GetTonKho(context.Background(), &pb.GetTonKhoRequest{Id: int32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, res.GetTonKho())
}
