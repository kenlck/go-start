package router

import (
	"crm_backend/handler"
	"crm_backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Hello)

	user := app.Group("/users")
	user.Get("/", middleware.Protected(), handler.GetAllUsers)
	user.Post("/create", handler.CreateUser)

	app.Post("/login", handler.Login)
}
