package routes

import (
	"qms/handler"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	bookHandler handler.BookHandlerInterface
	authHandler handler.AuthHandlerInterface
}

func NewRoutes(bookHandler handler.BookHandlerInterface, authHandler handler.AuthHandlerInterface) *Routes {
	return &Routes{
		bookHandler: bookHandler,
		authHandler: authHandler,
	}
}

func (r *Routes) Setup(app *fiber.App) {

	// auth
	app.Post("/register", r.authHandler.Register)
	app.Post("/login", r.authHandler.Login)
	app.Get("/user", r.authHandler.User)
	app.Post("/logout", r.authHandler.Logout)

	//check health server
	app.Get("/", r.bookHandler.HealthCheck)

	// booking
	app.Post("/book/create", r.bookHandler.CreateBook)
	app.Get("/bank", r.bookHandler.GetBank)
	app.Delete("/book/selesai/:id", r.bookHandler.DeleteBook)
	app.Get("/bank/detail/:id", r.bookHandler.GetBankDetailById)
	app.Get("/book/:id/", r.bookHandler.GetBookByUserId)
	app.Put("/book/selesai/:id", r.bookHandler.UpdateStatus)
	app.Get("/book/detail/:id/", r.bookHandler.GetBookById)

}
