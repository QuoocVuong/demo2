package models

type KhoHang struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	TenKho string `json:"ten_kho"`
}
