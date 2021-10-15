package database

import (
	"qms/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	//for Production
	// dsn := os.Getenv("CONNSTRING")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// For Online DB Development
	dsn := "sql6441433:K7cwfKcqVs@tcp(sql6.freesqldatabase.com:3306)/sql6441433"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	// For LocalDB Developmen
	// dsn := "user:password@tcp(localhost:3306)/database"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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