package main

import (
	"fmt"
	"log"

	"api-gateway/handlers"
	pbInventory "github.com/Q.Vuong/demo2/proto" // Import package proto của inventory service
	pbProduct "github.com/Q.Vuong/demo2/proto"   // Import package proto của product service
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Kết nối đến inventory-service
	connInventory, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Không thể kết nối đến inventory-service: %v", err)
	}
	defer connInventory.Close()
	inventoryClient := pbInventory.NewInventoryServiceClient(connInventory)

	// Kết nối đến product-catalog-service
	connProduct, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Không thể kết nối đến product-catalog-service: %v", err)
	}
	defer connProduct.Close()
	productClient := pbProduct.NewProductServiceClient(connProduct)

	// Khởi tạo handlers
	inventoryHandler := &handlers.InventoryHandler{InventoryClient: inventoryClient}
	productHandler := &handlers.ProductHandler{ProductClient: productClient}

	r := gin.Default()

	// Inventory routes
	inventory := r.Group("/inventory")
	{
		inventory.GET("/", inventoryHandler.GetAllTonKhos)
		inventory.GET("/:id", inventoryHandler.GetTonKho)
		// ... các routes khác cho inventory
	}

	// Product routes
	product := r.Group("/products")
	{
		product.GET("/", productHandler.GetAllSanPhams)
		product.GET("/:id", productHandler.GetSanPham)
		// ... các routes khác cho product
	}

	fmt.Println("API Gateway listening on port 8080")
	r.Run(":8080")
}
