package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"product-catalog-service/models"
	"product-catalog-service/repository"
)

type ProductHandler struct {
	// ... Các repository khác ...
	InventoryClient pb.InventoryServiceClient
}

var sanPhamRepository repository.SanPhamRepository

func (h *ProductHandler) GetProduct(c *gin.Context) {
	// ... (Lấy thông tin sản phẩm từ database) ...

	// Gọi gRPC đến inventory-service để lấy số lượng tồn kho
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("Error connecting to inventory service:", err)
		// ... Xử lý lỗi kết nối
	}
	defer conn.Close()

	client := inventory.NewInventoryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	inventoryRes, err := client.GetInventory(ctx, &inventory.GetInventoryRequest{ProductID: sanPham.ID})
	if err != nil {
		fmt.Println("Error calling GetInventory:", err)
		// ... Xử lý lỗi gọi gRPC
	}

	// ... (Kết hợp thông tin sản phẩm và tồn kho) ...

	c.JSON(http.StatusOK, gin.H{
		"product":  sanPham,
		"quantity": inventoryRes.GetQuantity(),
	})
}

type ProductHandler struct {
	// ... Các repository khác ...
	InventoryClient pb.InventoryServiceClient
}

func InitProductHandler(nhomHangRepo repository.NhomHangRepository /* ... Các repository khác, */, inventoryClient pb.InventoryServiceClient) {
	// ...
	productHandler = &ProductHandler{
		// ...
		InventoryClient: inventoryClient,
	}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	// ... (Lấy thông tin sản phẩm từ database) ...

	// Gọi gRPC đến inventory-service để lấy số lượng tồn kho
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	inventoryRes, err := h.InventoryClient.GetInventory(ctx, &pb.GetInventoryRequest{ProductID: int32(sanPham.ID)})
	if err != nil {
		fmt.Println("Error calling GetInventory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi lấy thông tin tồn kho"})
		return
	}

	if inventoryRes.GetError() != "" {
		// Xử lý lỗi từ inventory-service
		c.JSON(http.StatusInternalServerError, gin.H{"error": inventoryRes.GetError()})
		return
	}

	// ... (Kết hợp thông tin sản phẩm và tồn kho) ...

	c.JSON(http.StatusOK, gin.H{
		"product":  sanPham,
		"quantity": inventoryRes.GetQuantity(),
	})
}

func InitSanPhamHandler(repo repository.SanPhamRepository) {
	sanPhamRepository = repo
}

func GetAllSanPhams(c *gin.Context) {
	sanPhams, err := sanPhamRepository.GetAllSanPham()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sanPhams)
}

func GetSanPhamByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	sanPham, err := sanPhamRepository.GetSanPhamByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sản phẩm không tồn tại"})
		return
	}
	c.JSON(http.StatusOK, sanPham)
}

func CreateSanPham(c *gin.Context) {
	var sanPham models.SanPham
	if err := c.ShouldBindJSON(&sanPham); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sanPhamRepository.CreateSanPham(&sanPham); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sanPham)
}

func UpdateSanPham(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var sanPham models.SanPham
	if err := c.ShouldBindJSON(&sanPham); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sanPhamRepository.UpdateSanPham(uint(id), &sanPham); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sanPham)
}

func DeleteSanPham(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := sanPhamRepository.DeleteSanPham(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xoá thành công"})
}
