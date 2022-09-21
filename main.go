package main

import (
	"crm_backend/config"
	"crm_backend/database"
	_ "crm_backend/docs"
	"crm_backend/router"

	"github.com/gofiber/fiber/v2"
)

// @title       CRM Backend
// @version     1.0
// @description This is a template for golang backend with fiber framework.
// @contact.name  Ken Low
// @contact.email kenlck1990@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath      /
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
