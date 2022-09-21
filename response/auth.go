package response

import "github.com/gofiber/fiber/v2"

func MissingJWTResponseBody() *fiber.Map {
	return &fiber.Map{
		"message": "Missing or malformed JWT",
	}
}

func InvalidJWTResponseBody() *fiber.Map {
	return &fiber.Map{
		"message": "Invalid or expired JWT",
	}
}

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
