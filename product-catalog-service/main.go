package main

import (
	"fmt"
	_ "fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"product-catalog-service/handlers" // Import handlers từ package của bạn
	"product-catalog-service/models"   // Import models của bạn
	pb "product-catalog-service/proto" // Import package proto của product service
	"product-catalog-service/repository"

	pb "github.com/Q.Vuong/demo2/proto" // Thay thế bằng đường dẫn thực tế
)

var db *gorm.DB

func main() {

	// Thiết lập cơ sở dữ liệu (riêng cho dịch vụ này - nên dùng một cơ sở dữ liệu khác nếu có thể)
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối đến cơ sở dữ liệu:", err)
	}

	err = db.AutoMigrate(&models.NhomHang{}, &models.SanPham{}, &models.MucTuHang{})
	if err != nil {
		log.Fatal("Không thể migrate schema:", err)
	}
	// Khởi tạo gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Lỗi khi listen: %v", err)
	}
	s := grpc.NewServer() // Tạo gRPC server

	// Đăng ký ProductService
	pb.RegisterProductServiceServer(s, &handlers.ProductHandler{
		NhomHangRepository:  nhomHangRepo,
		SanPhamRepository:   sanPhamRepo,
		MucTuHangRepository: mucTuHangRepo,
	})

	fmt.Println("Product Catalog Service listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Lỗi khi serve: %v", err)
	}

	//s := grpc.NewServer()
	//
	//defer conn.Close()
	//var inventoryClient := pb.NewInventoryServiceClient(conn)
	// Khởi tạo repositories
	nhomHangRepo := repository.NewGormNhomHangRepository(db)
	sanPhamRepo := repository.NewGormSanPhamRepository(db)
	mucTuHangRepo := repository.NewGormMucTuHangRepository(db)

	// Khởi tạo handlers với inventoryClient
	handlers.InitProductHandler(nhomHangRepo, sanPhamRepo, mucTuHangRepo, inventoryClient)

	// Khởi tạo handlers
	handlers.InitNhomHangHandler(nhomHangRepo)
	handlers.InitSanPhamHandler(sanPhamRepo)
	handlers.InitMucTuHangHandler(mucTuHangRepo)
	r := gin.Default()

	nhomHangRoutes(r)  // Các routes liên quan đến NhomHang
	sanPhamRoutes(r)   // Các routes liên quan đến SanPham
	mucTuHangRoutes(r) // Các routes liên quan đến MucTuHang

	r.Run(":8081") // Chạy trên một cổng khác với các dịch vụ khác
}

func sanPhamRoutes(r *gin.Engine) {
	sanPham := r.Group("/san-pham")
	{
		sanPham.GET("/", handlers.GetAllSanPhams)
		sanPham.GET("/:id", handlers.GetSanPhamByID)
		sanPham.POST("/", handlers.CreateSanPham)
		sanPham.PUT("/:id", handlers.UpdateSanPham)
		sanPham.DELETE("/:id", handlers.DeleteSanPham)
	}
}

func mucTuHangRoutes(r *gin.Engine) {
	mucTuHang := r.Group("/muc-tu-hang")
	{
		mucTuHang.GET("/", handlers.GetAllMucTuHangs)
		mucTuHang.GET("/:id", handlers.GetMucTuHangByID)
		mucTuHang.POST("/", handlers.CreateMucTuHang)
		mucTuHang.PUT("/:id", handlers.UpdateMucTuHang)
		mucTuHang.DELETE("/:id", handlers.DeleteMucTuHang)
	}
}
