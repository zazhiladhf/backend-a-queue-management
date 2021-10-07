package repository

import (
	"github.com/slonob0y/qms/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

type BookRepoInterface interface {
	CreateBook(book models.SlotBooking) (models.SlotBooking, error)
	FindAllBank() ([]models.Bank, error)
	DeleteBook(status string) error
	GetBankById(id string) (models.Bank, error)
}

func(r *BookRepository) CreateBook(book models.SlotBooking) (models.SlotBooking, error) {
	trx := r.db.Create(&book)

	return book, trx.Error
}

func(r *BookRepository) FindAllBank() ([]models.Bank, error) {
	var banks []models.Bank
	findResult := r.db.Find(&banks)

	return banks, findResult.Error
}

func(r *BookRepository) DeleteBook(status string) error {
	var book models.SlotBooking
	findResult := r.db.Unscoped().Where("status = ?", status).Delete(&book)
	if findResult.Error != nil {
		return findResult.Error
	}

	if findResult.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	
	return nil
}

func(r *BookRepository) GetBankById(id string) (models.Bank, error) {
	var bank models.Bank
	findResult := r.db.Where("id = ?", id).First(&bank)
	return bank, findResult.Error
}