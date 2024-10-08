package models

type NhomHang struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TenNhom string `json:"ten_nhom"`
}
