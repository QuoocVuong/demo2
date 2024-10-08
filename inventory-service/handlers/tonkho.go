package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"models"
	//pb "inventory-service/proto" // Import package proto của inventory service
	pb "github.com/Q.Vuong/demo2/proto" // Thay thế bằng đường dẫn thực tế
	"repository"
)

type InventoryHandler struct {
	pb.UnimplementedInventoryServiceServer // Nên implement interface này để tránh lỗi khi thêm phương thức mới vào proto
	TonKhoRepository                       repository.TonKhoRepository
}

func (h *InventoryHandler) GetInventory(ctx context.Context, req *pb.GetInventoryRequest) (*pb.GetInventoryResponse, error) {
	fmt.Println("GetInventory called from gRPC client")
	tonKho, err := h.TonKhoRepository.GetTonKhoByProductID(req.GetProductID())
	if err != nil {
		return nil, fmt.Errorf("Lỗi khi lấy thông tin tồn kho: %v", err)
	}

	return &pb.GetInventoryResponse{
		ProductID: tonKho.ProductID,
		Quantity:  tonKho.Quantity,
	}, nil
}

// Tạo một biến toàn cục cho repository
var tonKhoRepository repository.TonKhoRepository

// Hàm khởi tạo để thiết lập repository
func InitTonKhoRepository(repo repository.TonKhoRepository) {
	tonKhoRepository = repo
}

// Lấy tất cả bản ghi tồn kho
func GetTonKhos(c *gin.Context) {
	tonKhos, err := tonKhoRepository.GetAllTonKho()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tonKhos)
}

// Lấy một bản ghi tồn kho theo ID
func GetTonKhoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tonKho, err := tonKhoRepository.GetTonKhoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tồn kho không tồn tại"})
		return
	}
	c.JSON(http.StatusOK, tonKho)
}

// Tạo một bản ghi tồn kho mới
func CreateTonKho(c *gin.Context) {
	var tonKho models.TonKho
	if err := c.ShouldBindJSON(&tonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tonKhoRepository.CreateTonKho(&tonKho); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tonKho)
}

// Cập nhật một bản ghi tồn kho
func UpdateTonKho(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var tonKho models.TonKho
	if err := c.ShouldBindJSON(&tonKho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tonKhoRepository.UpdateTonKho(uint(id), &tonKho); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tonKho)
}

// Xoá một bản ghi tồn kho
func DeleteTonKho(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := tonKhoRepository.DeleteTonKho(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xoá thành công"})
}
