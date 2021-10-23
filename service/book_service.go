package service

import (
	"errors"
	"fmt"

	"qms/models"
	"qms/repository"
	"qms/utils"
)

type BookService struct {
	bookRepo repository.BookRepoInterface
}

func NewBookService(bookRepo repository.BookRepoInterface) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

type BookServiceInterface interface {
	CreateBook(book *models.SlotBooking) (data *models.SlotBooking, err error)
	GetBank() ([]models.Bank, error)
	DeleteBook(status string) error
	UpdateBookStatus(book models.SlotBooking, id string) (books models.SlotBooking, err error)
	GetBankById(id string) ([]models.SlotBooking, error)
	// JoinTable() error
	ValidateBookByDay() error
	ValidateBookByToday() error
	GetBookByUserId(id string) ([]models.SlotBooking, error)
	GetBookById(id string) (models.SlotBooking, error)
}

func (s *BookService) CreateBook(book *models.SlotBooking) (data *models.SlotBooking, err error) {
	date := utils.FormatGetDate()
	hour := utils.FormatGetHour()

	book.TanggalPelayanan = date
	book.JamPelayanan = hour

	result, err := s.bookRepo.CreateBook(book, date, hour)
	if err != nil {
		return nil, err
	}
	book.ID = result

	return book, err
}

func (s *BookService) GetBank() ([]models.Bank, error) {
	banks, err := s.bookRepo.FindAllBank()

	return banks, err
}

func (s *BookService) DeleteBook(id string) error {
	err := s.bookRepo.DeleteBook(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) UpdateBookStatus(book models.SlotBooking, id string) (books models.SlotBooking, err error) {
	// update book data
	err = s.bookRepo.UpdateBookStatus(book, id)
	if err != nil {
		return books, err
	}

	// select book
	books, err = s.bookRepo.FindByID(id)
	if err != nil {
		return books, err
	}

	fmt.Println("books", books)

	return books, nil
}

func (s *BookService) GetBankById(id string) ([]models.SlotBooking, error) {
	// join table
	// bank, err := s.bookRepo.JoinTable()
	// if err != nil {
	// 	return bank, err
	// }

	// get bank
	bank, err := s.bookRepo.GetBankById(id)
	if err != nil {
		return bank, err
	}

	return bank, nil
}

// func (s *BookService) JoinTable() error {
// 	err := s.bookRepo.JoinTable()

// 	return err
// }

func (s *BookService) ValidateBookByDay() error {
	date := utils.FormatGetDate()
	count, err := s.bookRepo.GetBookByDate(date)
	if err != nil {
		return err
	}

	if count > 10 {
		return errors.New("booking penuh")
	}

	return nil

}

func (s *BookService) ValidateBookByToday() error {
	date := utils.FormatGetToday()
	_, err := s.bookRepo.GetBookByDate(date)
	if err != nil {
		return err
	}

	return nil

}

func (s *BookService) GetBookByUserId(id string) ([]models.SlotBooking, error) {
	book, err := s.bookRepo.GetBookByUserId(id)

	if err != nil {
		return book, err
	}
	fmt.Println(book)

	return book, nil
}

func (s *BookService) GetBookById(id string) (models.SlotBooking, error) {
	book, err := s.bookRepo.GetBookById(id)

	if err != nil {
		return book, err
	}
	fmt.Println(book)

	return book, nil
}
