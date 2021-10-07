package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/slonob0y/qms/models"
	"github.com/slonob0y/qms/service"
	"gorm.io/gorm"
)

type BookHandler struct {
	bookService service.BookServiceInterface
}

func NewBookHandler(bookService service.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

type BookHandlerInterface interface {
	CreateBook(c *fiber.Ctx) error
	GetBank(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
	GetBankById(c *fiber.Ctx) error
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := models.SlotBooking{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	response, err := h.bookService.CreateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"msg":    "success update data",
		"result": response,
	})
	
}

func(h *BookHandler) GetBank(c *fiber.Ctx) error {
	banks, err := h.bookService.GetBank()

	if err != nil {
		c.Status(fiber.StatusForbidden).JSON(err)
	}

	return c.JSON(banks)
}

func(h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	status := c.Params("status")

	err := h.bookService.DeleteBook(status)
	if err != nil {
		var statusCode = fiber.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = fiber.StatusNotFound
		}

		return c.Status(statusCode).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"result": "success delete data",
	})
	
}

func(h *BookHandler) GetBankById(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.bookService.GetBankById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		// "msg":    "success update data",
		"result": response,
	})
	
}