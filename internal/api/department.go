package api

import (
	"context"
	"net/http"
	"time"

	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/golang-jwt/jwt/v5"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/internal/middleware"
	"github.com/gofiber/fiber/v2"
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
	user.Get("/", da.GetDepartment)
	// user.Get("/", da.Index)
	user.Patch("/:id", da.UpdateDepartment)
	user.Delete("/:id", da.DeleteDepartment)
}

// func (da departmentApi) Index(ctx *fiber.Ctx) error {
// 	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
// 	defer cancel()

// 	res, code, err := da.departmentService.Index(c)

// 	if err != nil {
// 		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
// 	}

// 	return ctx.Status(code).JSON(res)
// }

func (da departmentApi) GetDepartment(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	res, code, err := da.departmentService.GetDepartment(ctx.Context(), ctx.Params("id"), userId)

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
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// KERJAIN DISINI BANG

	return ctx.Status(400).JSON(fiber.Map{})
}
