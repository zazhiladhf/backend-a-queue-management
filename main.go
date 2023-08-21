package main

import (
	"qms/database"
	"qms/handler"
	"qms/repository"
	"qms/routes"
	"qms/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//set up connection
	DB := database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// app.Get("/", Book)
	// app.Post("/", PostBook)

	// book services
	bookRepository := repository.NewBookRepository(DB)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// auth services
	var authHandler = &handler.AuthHandler{}

	routes := routes.NewRoutes(bookHandler, authHandler)
	routes.Setup(app)

	app.Listen(":8080")
	// app.Listen(":" + os.Getenv("APP_PORT"))
}
