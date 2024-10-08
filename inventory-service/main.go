package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"inventory-service/handlers"
	"inventory-service/models"
	"inventory-service/repository"
	//"google.golang.org/grpc"
	//pb "inventory-service/proto" // Import package proto của inventory service
	pb "github.com/Q.Vuong/demo2/proto" // Thay thế bằng đường dẫn thực tế
)

var (
	db          *gorm.DB
	khoHangRepo repository.KhoHangRepository
	tonKhoRepo  repository.TonKhoRepository

func main() {

	// Thiết lập kết nối CSDL cho inventory service (database riêng)
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo2?charset=utf8mb4&parseTime=True&loc=Local" // Thay đổi thông tin kết nối nếu cần
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Lỗi kết nối CSDL:", err)
	}

	// Migrate schema cho inventory service
	err = db.AutoMigrate(&models.KhoHang{}, &models.TonKho{})
	if err != nil {
		log.Fatal("Lỗi migrate schema:", err)
	}

	// Khởi tạo repository
	khoHangRepo = repository.NewKhoHangRepository(db)
	tonKhoRepo = repository.NewTonKhoRepository(db)

	//handlers.Init(db, khoHangRepo , tonKhoRepo) // Khởi tạo handlers với kết nối DB

	// Khởi tạo handler và truyền repository
	//handlers.InitKhoHangHandler(khoHangRepo)
	if err := handlers.InitKhoHangHandler(khoHangRepo); err != nil {
		log.Fatal("Lỗi khi khởi tạo KhoHangHandler:", err)
	}

	// Khởi tạo gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Lỗi khi listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &handlers.InventoryHandler{
		TonKhoRepository: tonKhoRepo,
	})
	fmt.Println("Inventory Service listening on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Lỗi khi serve: %v", err)
	}

	r := gin.Default()

	// Routes cho KhoHang
	khoHangRoutes(r)

	// Routes cho TonKho
	tonKhoRoutes(r)

	r.Run(":8082") // Chạy trên cổng 8082 (khác với product-catalog-service)
}

func khoHangRoutes(r *gin.Engine) {
	khoHang := r.Group("/kho-hang")
	{
		khoHang.GET("/", handlers.GetKhoHangs)
		khoHang.GET("/:id", handlers.GetKhoHang)
		khoHang.POST("/", handlers.CreateKhoHang)
		khoHang.PUT("/:id", handlers.UpdateKhoHang)
		khoHang.DELETE("/:id", handlers.DeleteKhoHang)
	}
}

//func tonKhoRoutes(r *gin.Engine) {
//	r.GET("/ton-kho", handlers.GetTonKhos)
//	r.GET("/ton-kho/:id", handlers.GetTonKho)
//	r.POST("/ton-kho", handlers.CreateTonKho)
//	r.PUT("/ton-kho/:id", handlers.UpdateTonKho)
//	r.DELETE("/ton-kho/:id", handlers.DeleteTonKho)
//}
func tonKhoRoutes(r *gin.Engine) {
	tonKho := r.Group("/ton-kho")
	{
		tonKho.GET("/", handlers.GetTonKhos)
		tonKho.GET("/:id", handlers.GetTonKho)
		tonKho.POST("/", handlers.CreateTonKho)
		tonKho.PUT("/:id", handlers.UpdateTonKho)
		tonKho.DELETE("/:id", handlers.DeleteTonKho)
	}
	}
}