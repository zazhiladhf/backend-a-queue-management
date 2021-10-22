package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	ID        uint   `gorm:"column:id;primarykey" json:"id_bank_tujuan"`
	Nama      string `gorm:"unique;not null;index" json:"nama_bank"`
	Alamat    string `gorm:"not null;index" json:"alamat"`
	Kapasitas int    `json:"kapasitas"`
	// SlotBookings []SlotBooking `gorm:"many2many:bank_details;"`
}

func InsertBank(db *gorm.DB) {
	db.Model(&Bank{}).Create([]map[string]interface{}{
		{"Nama": "BANK KCP SOREANG", "Alamat": "Jl.Soreang No.180", "Kapasitas": 20},
		{"Nama": "BANK KCP Banjaran", "Alamat": "Jl.Banjaran No.181", "Kapasitas": 20},
	})
}
