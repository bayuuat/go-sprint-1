package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type departmentApi struct {
	departmentService domain.DepartmentService
}

func NewDepartment(app *fiber.App,
	departmentService domain.DepartmentService) {

	da := departmentApi{
		departmentService: departmentService,
	}

	user := app.Group("/v1/department")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreateDepartment)
	user.Get("/", da.GetDepartments)
	user.Patch("/:id", da.UpdateDepartment)
	user.Delete("/:id", da.DeleteDepartment)
}

func (da departmentApi) GetDepartments(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	filter := dto.DepartmentFilter{
		Limit:  5,
		Offset: 0,
		Name:   ctx.Query("name", ""),
		UserId: userId,
	}

	if limit, err := strconv.Atoi(ctx.Query("limit", "5")); err == nil && limit > 0 {
		filter.Limit = limit
	}
	if offset, err := strconv.Atoi(ctx.Query("offset", "0")); err == nil && offset >= 0 {
		filter.Offset = offset
	}

	res, code, err := da.departmentService.GetDepartmentsWithFilter(ctx.Context(), filter)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da departmentApi) CreateDepartment(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}

func (da departmentApi) UpdateDepartment(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// Get user_id claims
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var req dto.DepartmentReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, code, err := da.departmentService.PatchDepartment(c, req, ctx.Params("id"), userId)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da departmentApi) DeleteDepartment(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["id"].(string)

	res, code, err := da.departmentService.DeleteDepartment(c, user_id, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
