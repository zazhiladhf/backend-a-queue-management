package models

type SlotBooking struct {
	ID               int    `gorm:"column:id" json:"id_booking"`
	TanggalPelayanan string `json:"tanggal_pelayanan"`
	JamPelayanan     string `json:"jam_pelayanan"`
	KeperluanLayanan int    `json:"keperluan_layanan"`
	Status           string `json:"status"`
	BankID           uint   `gorm:"column:bank_id;unique" json:"id_bank_tujuan" validate:"required"`
	UserID           uint   `json:"id_user"`
}
