package api

import (
	"context"
	"time"
	"net/http"
	"strconv"
	"fmt"
	"log"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/middleware"
	"github.com/bayuuat/go-sprint-1/internal/utils"
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

	var req dto.EmployeeReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.NewErrorResponse("Invalid request:" + err.Error()))
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid request"))
	}

	id, msg, err := da.employeeService.CreateEmployee(ctx.Context(), req, req.Name)
	if err != nil {
		log.Println(fmt.Sprintf("Error create employee: %v", err))
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse(strconv.Itoa(msg) + err.Error() + " disini"))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.NewSuccessCreateResponse((strconv.Itoa(msg)), id))
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
