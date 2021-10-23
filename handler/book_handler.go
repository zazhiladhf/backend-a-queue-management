package handler

import (
	"errors"

	"qms/models"
	"qms/service"
	"qms/utils"

	"github.com/gofiber/fiber/v2"
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
	HealthCheck(c *fiber.Ctx) error
	CreateBook(c *fiber.Ctx) error
	GetBank(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
	GetBankDetailById(c *fiber.Ctx) error
	GetBookById(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	GetBookByUserId(c *fiber.Ctx) error
}

func (h *BookHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"server":  true,
		"message": "Server UP Capt 🚀",
	})
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := &models.SlotBooking{}
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()
	// book.BankID =
	if err := validate.Struct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	err = h.bookService.ValidateBookByDay()
	if err != nil {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  201,
			"message": err.Error(),
		})
	}

	// if book.ID > 7 {
	// 	c.Status(fiber.StatusCreated)
	// 	return c.JSON(fiber.Map{
	// 		"status":  201,
	// 		"message": "booking penuh",
	// 		"data":    book,
	// 	})
	// }

	response, err := h.bookService.CreateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  true,
			"msg":    err.Error(),
			"status": 500,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  200,
		"message": "berhasil",
		"data":    response,
	})

}

func (h *BookHandler) GetBank(c *fiber.Ctx) error {
	banks, err := h.bookService.GetBank()

	if err != nil {
		c.Status(fiber.StatusForbidden).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  200,
		"message": "berhasil",
		"data":    banks,
	})
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.bookService.DeleteBook(id)
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "berhasil",
	})

}

func (h *BookHandler) GetBankDetailById(c *fiber.Ctx) error {
	id := c.Params("id")
	// book := &models.SlotBooking{}
	// err := c.BodyParser(book)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// err := h.bookService.JoinTable()
	// if err != nil {
	// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	// 		"status":  201,
	// 		"message": err.Error(),
	// 	})
	// }

	// err = h.bookService.ValidateBookByToday()
	// if err != nil {
	// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	// 		"status":  201,
	// 		"message": err.Error(),
	// 	})
	// }

	response, err := h.bookService.GetBankById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    err.Error(),
			"status": 404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "berhasil",
		"result":  response,
		"status":  200,
	})

}

func (r *BookHandler) GetBookByUserId(c *fiber.Ctx) error {
	// status := c.Query("status")
	// id := c.Query("id")
	// status := c.Params("status")
	id := c.Params("id")

	// var book models.SlotBooking

	// if err := c.BodyParser(&book); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// fmt.Println("hand", book)

	// _, err := r.bookService.UpdateBookStatus(book, status)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	response, err := r.bookService.GetBookByUserId(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		// "msg":    "success update data",
		"result": response,
	})

}

func (m *BookHandler) UpdateStatus(c *fiber.Ctx) error {
	book := models.SlotBooking{}
	id := c.Params("id")
	// var mysqlErr *mysql.MySQLError

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	response, err := m.bookService.UpdateBookStatus(book, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    "success update data",
		"result": response,
	})
}

func (r *BookHandler) GetBookById(c *fiber.Ctx) error {
	// status := c.Query("status")
	// id := c.Query("id")
	// status := c.Params("status")
	id := c.Params("id")

	response, err := r.bookService.GetBookById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		// "msg":    "success update data",
		"result": response,
	})

}
