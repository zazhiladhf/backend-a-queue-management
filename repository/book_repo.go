package repository

import (
	"fmt"

	"qms/models"

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
	CreateBook(book *models.SlotBooking, date, hour string) (uint, error)
	FindAllBank() ([]models.Bank, error)
	FindByStatus(status string) (models.SlotBooking, error)
	DeleteBook(id string) error
	UpdateBookStatus(book models.SlotBooking, status string) error
	GetBankById(id string) (models.Bank, error)
	JoinTable() error
	GetBookByDate(date string) (uint, error)
	GetBookByUserId(id string) ([]models.SlotBooking, error)
	GetBookById(id string) (models.SlotBooking, error)
}

func (r *BookRepository) CreateBook(book *models.SlotBooking, date, hour string) (uint, error) {
	result := r.db.Create(book)
	if result.Error != nil {
		return book.ID, result.Error
	}

	return book.ID, nil
}

// func(r *BookRepository) CreateBook(book *models.SlotBooking) error {
// 	query := `INSERT INTO slot_bookings (id_booking, tanggal_pelayanan, jam_pelayanan, keperluan_layanan, status, id_bank_tujuan, id_user)
// 				VALUES ($1, $2, $3, $4, $5, $6, $7)`
// 	_, err := r.db.Exec(query, book.ID, book.TanggalPelayanan, book.JamPelayanan, book.KeperluanLayanan, book.Status, book.BankID, book.UserID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func(r *BookRepository) CreateFullBook(book models.SlotBooking) (models.SlotBooking, error) {
// 	trx := r.db.Create(&book)

// 	return book, trx.Error
// }

func (r *BookRepository) FindAllBank() ([]models.Bank, error) {
	banks := []models.Bank{}
	query := `SELECT id, nama, alamat FROM banks`
	// findResult := r.db.Find(&banks)

	err := r.db.Raw(query).Scan(&banks)
	if err != nil {
		return banks, err.Error
	}

	return banks, nil
}

func (r *BookRepository) FindByStatus(status string) (models.SlotBooking, error) {
	var book models.SlotBooking
	findResult := r.db.Where("status = ?", status).First(&book)
	return book, findResult.Error
}

func (r *BookRepository) DeleteBook(id string) error {
	var book models.SlotBooking
	findResult := r.db.Where("id = ?", id).Delete(&book)
	// if err != nil {
	// 	return err
	// }
	if findResult.Error != nil {
		return findResult.Error
	}

	if findResult.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *BookRepository) UpdateBookStatus(book models.SlotBooking, status string) error {
	query := `UPDATE slot_bookings SET status = "done"`
	result := r.db.Exec(query, book.Status, status)
	// fmt.Println("result",result)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// trx := r.DB.Where("slug = ?", slug).Update("title", "edit")
	// // trx := r.DB.Model(&model.Movie).Update(&movie)

	return nil
}

func (r *BookRepository) GetBankById(id string) (models.Bank, error) {
	var bank models.Bank

	findResult := r.db.Where("id = ?", id).First(&bank)
	return bank, findResult.Error
}

func (r *BookRepository) JoinTable() error {
	var bank models.Bank
	query := `SELECT banks.id, nama, alamat, tanggal_pelayanan, slot_bookings.id, jam_pelayanan FROM banks INNER JOIN slot_bookings ON banks.id=slot_bookings.bank_id`

	err := r.db.Raw(query).Scan(&bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) GetBookByDate(date string) (uint, error) {
	var count uint
	query := `SELECT COUNT(*) FROM slot_bookings where tanggal_pelayanan = ?`

	err := r.db.Raw(query, date).Scan(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}

func (r *BookRepository) GetBookByUserId(id string) ([]models.SlotBooking, error) {
	// book := models.SlotBooking{}
	var book []models.SlotBooking

	query := `SELECT id, tanggal_pelayanan, jam_pelayanan, keperluan_layanan, status, bank_id, user_id FROM slot_bookings WHERE user_id = ?`

	err := r.db.Raw(query, id).Scan(&book).Error
	if err != nil {
		return book, err
	}

	// if book.ID == 0 {
	// 	return book, gorm.ErrRecordNotFound
	// }
	// findResult := r.db.Limit(2).Where("user_id = ?", id).First(&book)
	fmt.Println("book", book)

	return book, nil
}

func (r *BookRepository) GetBookById(id string) (models.SlotBooking, error) {
	var book models.SlotBooking

	query := `SELECT id, tanggal_pelayanan, jam_pelayanan, keperluan_layanan, status, bank_id, user_id FROM slot_bookings WHERE id = ?`

	err := r.db.Raw(query, id).Scan(&book).Error
	if err != nil {
		return book, err
	}

	fmt.Println("book", book)

	return book, nil
}
