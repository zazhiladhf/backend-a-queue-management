package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	ID       uint   `gorm:"primaryKey" json:"id_bank"`
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	BookingSlot uint ` json:"booking_slot"`
	SlotBooking []*SlotBooking
}

func InsertBank(db *gorm.DB)  {
	db.Model(&Bank{}).Create([]map[string]interface{}{
		{"Nama": "BANK KCP SOREANG", "Alamat": "Jl.Soreang No.180", "BookingSlot": 40},
  		{"Nama": "BANK KCP Banjaran", "Alamat": "Jl.Banjaran No.181", "BookingSlot": 40},
	})
}