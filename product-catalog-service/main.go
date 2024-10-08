package main

import (

	_ "fmt"
	"log"


	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"handlers" // Import handlers từ package của bạn
	"models"
	"repository"// Import models của bạn

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
	// Khởi tạo gRPC connection đến inventory-service
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Không thể kết nối đến inventory-service: %v", err)
	}
	defer conn.Close()
	var inventoryClient := pb.NewInventoryServiceClient(conn)
	// Khởi tạo repositories
	nhomHangRepo := repository.NewGormNhomHangRepository(db)
	sanPhamRepo := repository.NewGormSanPhamRepository(db)
	mucTuHangRepo := repository.NewGormMucTuHangRepository(db)

	// Khởi tạo handlers với inventoryClient
	handlers.InitProductHandler(nhomHangRepo,sanPhamRepo,mucTuHangRepo /* ... các repository khác, */ inventoryClient)

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
