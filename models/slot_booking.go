package models

type SlotBooking struct {
	ID               uint   `gorm:"column:id" json:"id_booking"`
	TanggalPelayanan string `json:"tanggal_pelayanan"`
	JamPelayanan     string `json:"jam_pelayanan"`
	KeperluanLayanan int    `json:"keperluan_layanan"`
	Status           string `json:"status"`
	BankID           uint   `json:"id_bank_tujuan"`
	UserID           uint   `json:"id_user"`
}
