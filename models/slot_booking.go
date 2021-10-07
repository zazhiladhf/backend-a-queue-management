package models

import "time"

type SlotBooking struct {
	ID       			uint   `gorm:"primaryKey" json:"id_slot_booking"`
	TanggalPelayanan 	time.Time
	JamPelayanan 		time.Time
	KeperluanLayanan	string `gorm:"primaryKey" json:"keperluan_layanan"`
	Status 				string `gorm:"primaryKey" json:"status"`
	BankID 				uint `json:"id_bank"`
	UserID 				uint `json:"id_user"`
}