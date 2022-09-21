package response

import "github.com/gofiber/fiber/v2"

func LoginFailResponseBody() *fiber.Map {
	return &fiber.Map{
		"message": "Invalid email or password",
	}
}

func LoginSuccessResponseBody(t string) *fiber.Map {
	return &fiber.Map{
		"message": "Login successful",
		"token":   t,
	}
}
