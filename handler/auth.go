package handler

import (
	"crm_backend/config"
	"crm_backend/database"
	"crm_backend/model"
	"crm_backend/response"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(ctx *fiber.Ctx) error {
	type incomingUser struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	var body incomingUser

	if err := ctx.BodyParser(&body); err != nil {
		fmt.Println(err.Error())
		return ctx.Status(401).JSON(response.LoginFailResponseBody())
	}

	user, err := getUserByEmail(body.Email)
	if err != nil {
		return ctx.Status(401).JSON(response.LoginFailResponseBody())
	}

	if !CheckPasswordHash(body.Password, user.Password) {
		return ctx.Status(401).JSON(response.LoginFailResponseBody())
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24 * 1).Unix() // 1 day
	t, err := token.SignedString([]byte(config.SECRET))
	if err != nil {
		return ctx.Status(401).JSON(response.LoginFailResponseBody())
	}

	return ctx.Status(200).JSON(response.LoginSuccessResponseBody(t))
}
