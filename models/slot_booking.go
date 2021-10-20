package models

type SlotBooking struct {
	ID               uint   `gorm:"column:id;primarykey" json:"id_booking"`
	TanggalPelayanan string `json:"tanggal_pelayanan"`
	JamPelayanan     string `json:"jam_pelayanan"`
	KeperluanLayanan int    `json:"keperluan_layanan"`
	Status           string `json:"status"`
	BankID           uint   `gorm:"column:bank_id;not null;index" json:"id_bank_tujuan"`
	UserID           int    `gorm:"column:user_id;not null;index" json:"id_user"`
	Bank             Bank
	User             User
	Banks            []Bank `gorm:"many2many:bank_details;"`
}
