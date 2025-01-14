package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/middleware"
	"github.com/bayuuat/go-sprint-1/internal/utils"
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
	user.Patch("/", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	})
	user.Patch("/:id?", da.UpdateDepartment)
	user.Delete("/", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	})
	user.Delete("/:id?", da.DeleteDepartment)
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

	var req dto.DepartmentReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid request:" + err.Error()))
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		var errMsg string
		for field, err := range fails {
			errMsg += field + ": " + err + "; "
		}
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("Validation error:  " + errMsg))
	}

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	res, _, err := da.departmentService.CreateDepartment(ctx.Context(), req, userId)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"departmentId": res.DepartmentId,
		"name":         res.Name,
	})
}

func (da departmentApi) UpdateDepartment(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	departmentId := ctx.Params("id")

	if departmentId == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	}

	if _, err := strconv.Atoi(departmentId); err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	}

	// Get user_id claims
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var req dto.UpdateDepartmentReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := da.departmentService.PatchDepartment(c, req, departmentId, userId)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da departmentApi) DeleteDepartment(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	}

	if _, err := strconv.Atoi(id); err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Not Found",
		})
	}

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["id"].(string)

	res, code, err := da.departmentService.DeleteDepartment(c, user_id, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
