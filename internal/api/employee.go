package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

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
	user.Patch("/:id?", da.UpdateEmployee)
	user.Delete("/:id?", da.DeleteEmployee)
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

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var req dto.EmployeeReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid request:" + err.Error()))
	}

	isIDEmployeeExists, err := da.employeeService.IsEmployeeIDExists(ctx.Context(), req.IdentityNumber, userId)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewErrorResponse("Internal server error"))
	}

	if isIDEmployeeExists {
		return ctx.Status(http.StatusConflict).JSON(dto.NewErrorResponse("Employee ID already exists"))
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if valid := middleware.ValidateUrl(req.EmployeeImageUri); !valid {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("invalid url"))
	}

	id, msg, err := da.employeeService.CreateEmployee(ctx.Context(), req, userId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse(strconv.Itoa(msg) + err.Error()))
	}

	return ctx.Status(201).JSON(dto.EmployeeData{
		IdentityNumber:   id.IdentityNumber,
		Name:             id.Name,
		EmployeeImageUri: id.EmployeeImageUri,
		Gender:           id.Gender,
		DepartmentID:     id.DepartmentID,
	})
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
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	if valid := middleware.ValidateUrl(req.EmployeeImageUri); !valid {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("invalid url"))
	}

	// invalid request
	employeePatch, err := req.Validate()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: domain.ErrBadRequest.Error()})
	}

	res, code, err := da.employeeService.PatchEmployee(c, req, ctx.Params("id"), userId, employeePatch)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da employeeApi) DeleteEmployee(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	fmt.Println()
	fmt.Println(id)

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["id"].(string)

	res, code, err := da.employeeService.DeleteEmployee(c, user_id, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
