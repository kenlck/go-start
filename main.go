package main

import (
	"crm_backend/config"
	"crm_backend/database"
	"crm_backend/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.LoadEnv()
	database.ConnectDB()

	router.SetupRoutes(app)

	err := app.Listen("localhost:8080")

	if err != nil {
		panic(err)
	}
}
