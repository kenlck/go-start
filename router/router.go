package router

import (
	"crm_backend/handler"
	"crm_backend/middleware"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Hello)

	v1 := app.Group("/v1")
	user := v1.Group("/users")
	user.Get("/", middleware.Protected(), handler.GetAllUsers)
	user.Post("/create", handler.CreateUser)

	app.Post("/login", handler.Login)

	app.Get("/swagger/*", swagger.HandlerDefault)

}
