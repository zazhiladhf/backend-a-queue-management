package database

import (
	"github.com/slonob0y/qms/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:74712331@/project"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = db

	models.InsertBank(db)

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Bank{})
	db.AutoMigrate(&models.SlotBooking{})

	return db
}
