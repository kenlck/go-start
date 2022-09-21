package main

import (
	"crm_backend/config"
	"crm_backend/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.LoadEnv()
	database.ConnectDB()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world!")
	})

	err := app.Listen("localhost:8080")

	if err != nil {
		panic(err)
	}
}
