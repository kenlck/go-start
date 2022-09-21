package handler

import (
	"crm_backend/database"
	"crm_backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	type incomingUser struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	db := database.DB

	var body incomingUser
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(400).JSON(fiber.Map{"status": "error"})
	}
	var user model.User
	db.Find(&user, "email = ?", body.Email)
	if CheckPasswordHash(body.Password, user.Password) {
		return ctx.Status(200).JSON(user)
	}

	return ctx.Status(401).JSON(fiber.Map{"status": "error"})
}
