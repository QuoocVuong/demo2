package models

type SanPham struct {
	ID                           uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	MaHang                       string   `json:"ma_hang"`
	TenMuc                       string   `json:"ten_muc"`
	NhomHangID                   uint     `json:"nhom_hang_id"`
	DonViDoDinh                  string   `json:"don_vi_do_dinh"`
	VoHieuHoa                    bool     `json:"vo_hieu_hoa"`
	ChoPhepKhoanThayThe          bool     `json:"cho_phep_khoan_thay_the"`
	DuyTriHangTonKho             bool     `json:"duy_tri_hang_ton_kho"`
	BaoQumCacMatHangTrongSanXuat bool     `json:"bao_qum_cac_mat_hang_trong_san_xuat"`
	CoPhieuMoDau                 string   `json:"co_phieu_mo_dau"`
	DinhGia                      float64  `json:"dinh_gia"`
	TyGiaBanHangTauChuan         float64  `json:"ty_gia_ban_hang_tau_chuan"`
	LaChiDinhTaiSan              bool     `json:"la_chi_dinh_tai_san"`
	NhomHang                     NhomHang `gorm:"foreignKey:NhomHangID" json:"nhom_hang"` // Add this for eager loading
}
