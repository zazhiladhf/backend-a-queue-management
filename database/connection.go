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

	// For Deployment Kominfo
	dsn := "backend-a-queue-management:dKq8hLYhbWF8sMXd@tcp(35.201.240.187:3306)/backend-a-queue-management"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// For Online DB Development
	// dsn := "sql6441433:K7cwfKcqVs@tcp(sql6.freesqldatabase.com:3306)/sql6441433"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// For LocalDB Developmen
	// dsn := "root:74712331@tcp(localhost:3306)/project"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Bank{})
	db.AutoMigrate(&models.SlotBooking{})

	models.InsertBank(db)

	return db
}
