package handler

import (
	"crm_backend/database"
	"crm_backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func GetAllUsers(ctx *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)
	return ctx.Status(200).JSON(users)
}

func CreateUser(ctx *fiber.Ctx) error {
	type incomingUser struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	db := database.DB

	var body incomingUser
	err := ctx.BodyParser(&body)

	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(400).JSON(fiber.Map{"status": "error"})
	}

	hash, err := hashPassword(body.Password)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error"})

	}
	user := model.User{
		Email:    body.Email,
		Password: hash,
		Name:     body.Name,
	}
	err = db.Create(&user).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Unable to create user"})
	}

	return ctx.Status(200).JSON(user)
}
