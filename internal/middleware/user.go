package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(authService domain.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.Split(c.Get("Authorization"), " ")
		if len(token) < 2 {
			return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "Sorry, the token you entered is invalid. Please check your token and try again or contact customer support for further assistance. Thank you."})
		}
		user, err := authService.Validate(context.Background(), token[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{Message: "Sorry, the token you entered is invalid. Please check your token and try again or contact customer support for further assistance. Thank you."})
		}

		c.Locals("x-user", user)
		return c.Next()
	}
}
