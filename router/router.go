package router

import (
	"crm_backend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Hello)

	user := app.Group("/users")
	user.Get("/", handler.GetAllUsers)
	user.Post("/create", handler.CreateUser)

	app.Post("/login", handler.Login)
}
