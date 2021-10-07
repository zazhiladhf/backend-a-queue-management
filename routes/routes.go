package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slonob0y/qms/handler"
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
	app.Post("/api/register", r.authHandler.Register)
	app.Post("/api/login", r.authHandler.Login)
	app.Get("/api/user", r.authHandler.User)
	app.Post("/api/logout", r.authHandler.Logout)

	// booking
	app.Post("/book/create", r.bookHandler.CreateBook)
	app.Get("/bank", r.bookHandler.GetBank)
	app.Delete("/book/:status", r.bookHandler.DeleteBook)
	app.Get("/bank/detail/:id", r.bookHandler.GetBankById)

}