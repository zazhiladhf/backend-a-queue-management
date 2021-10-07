package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/slonob0y/qms/database"
	"github.com/slonob0y/qms/handler"
	"github.com/slonob0y/qms/models"
	"github.com/slonob0y/qms/repository"
	"github.com/slonob0y/qms/routes"
	"github.com/slonob0y/qms/service"
)

func Book(c *fiber.Ctx) error {
	book := models.SlotBooking{
		ID:               1,
		TanggalPelayanan: time.Now(),
		JamPelayanan:     time.Now(),
		KeperluanLayanan: "transaksi",
		Status:           "done",
		BankID:           1,
		UserID:           1,
	}
	return c.JSON(book)
}

func PostBook(c *fiber.Ctx) error {
	book := &models.SlotBooking{}

	if err := c.BodyParser(book); err != nil {
		c.Status(403).JSON(err)
	}
	return c.JSON(book)
}

func main() {

	//set up connection
	DB := database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", Book)
	app.Post("/", PostBook)

	// book services
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// auth services
	var authHandler = &handler.AuthHandler{}

	routes := routes.NewRoutes(bookHandler, authHandler)
	routes.Setup(app)

	app.Listen(":8000")
}