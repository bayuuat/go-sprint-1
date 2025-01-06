package api

import (
	"context"
	"net/http"
	"time"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	authService domain.UserService
}

func NewUser(app *fiber.App, authHandler fiber.Handler,
	authService domain.UserService) {

	ha := authApi{
		authService: authService,
	}

	app.Post("/v1/auth", ha.authenticate)
}

func (a authApi) authenticate(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, code, err := a.authService.Authenticate(c, req)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
