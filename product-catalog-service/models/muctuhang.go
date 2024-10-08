package models

type MucTuHang struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	SanPhamID uint    `json:"san_pham_id"`
	TenMucTu  string  `json:"ten_muc_tu"`
	SanPham   SanPham `gorm:"foreignKey:SanPhamID" json:"san_pham"`
}
