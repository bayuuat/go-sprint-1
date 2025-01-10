package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/golang-jwt/jwt/v5"

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
	user.Patch("/:id", da.UpdateEmployee)
	user.Delete("/:id", da.DeleteEmployee)
}

func (da employeeApi) GetEmployee(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	limit := ctx.QueryInt("limit", 5)
	offset := ctx.QueryInt("offset", 0)
	name := ctx.Query("name", "")
	identityNumber := ctx.Query("identityNumber", "")
	gender := ctx.Query("gender", "")
	departmentId := ctx.Query("departmentId", "")

	filter := dto.EmployeeFilter{
		Limit:          limit,
		Offset:         offset,
		Name:           name,
		UserId:         userId,
		IdentityNumber: identityNumber,
		Gender:         gender,
		DepartmentID:   departmentId,
	}

	employees, err := da.employeeService.GetEmployees(ctx.Context(), filter)
	if err != nil {
		log.Println("Error getting employees:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	if len(employees) == 0 {
		return ctx.Status(200).JSON([]interface{}{})
	}

	return ctx.Status(200).JSON(employees)
}

func (da employeeApi) CreateEmployee(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}

func (da employeeApi) UpdateEmployee(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// Get user_id claims
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var req dto.EmployeeReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, code, err := da.employeeService.PatchEmployee(c, req, ctx.Params("id"), userId)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da employeeApi) DeleteEmployee(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["id"].(string)

	fmt.Print(user_id)

	res, code, err := da.employeeService.DeleteEmployee(c, user_id, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
