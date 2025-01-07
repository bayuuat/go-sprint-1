package api

import (
	"context"
	"time"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type employeeApi struct {
	employeeService domain.EmployeeService
}

func NewEmployee(app *fiber.App,
	employeeService domain.EmployeeService) {

	da := employeeApi{
		employeeService: employeeService,
	}

	user := app.Group("/v1/employee")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreateEmployee)
	user.Get("/", da.GetEmployee)
	user.Patch("/:id", da.UpdatedEmployee)
	user.Delete("/:id", da.DeleteEmployee)
}

func (da employeeApi) GetEmployee(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}

func (da employeeApi) CreateEmployee(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}

func (da employeeApi) UpdatedEmployee(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}

func (da employeeApi) DeleteEmployee(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}
