package service

import (
	"fmt"

	"github.com/slonob0y/qms/models"
	"github.com/slonob0y/qms/repository"
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
	CreateBook(book models.SlotBooking) (models.SlotBooking,error)
	GetBank() ([]models.Bank, error)
	DeleteBook(status string) error
	GetBankById(id string) (models.Bank, error)
}

func(s *BookService) CreateBook(book models.SlotBooking) (models.SlotBooking, error) {

	response, err := s.bookRepo.CreateBook(book)
	fmt.Println(response)

	return response, err
}

func(s *BookService) GetBank() ([]models.Bank, error) {
	banks, err := s.bookRepo.FindAllBank()

	return banks, err
}

func(s *BookService) DeleteBook(status string) error {
	err := s.bookRepo.DeleteBook(status)

	if err != nil {
		return err
	}

	return nil
}

func(s *BookService) GetBankById(id string) (models.Bank, error) {
	movies, err := s.bookRepo.GetBankById(id)

	return movies, err
}