package domain

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/dto"
)

type Department struct {
	DepartmentId string       `db:"department_id"`
	Name         string       `db:"name"`
	CreatedAt    sql.NullTime `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	UserId       string       `db:"user_id"`
}

type DepartmentRepository interface {
	Save(ctx context.Context, department *Department) error
	Update(ctx context.Context, department *Department) error
	FindAllWithFilter(ctx context.Context, filter *dto.DepartmentFilter) ([]Department, error)
	FindById(ctx context.Context, id string, userId string) (Department, error)
	HasEmployees(ctx context.Context, departmentId string) (bool, error)
	Delete(ctx context.Context, user_id string, id string) error
}

type DepartmentService interface {
	GetDepartmentsWithFilter(ctx context.Context, filter dto.DepartmentFilter) ([]dto.DepartmentData, int, error)
	CreateDepartment(ctx context.Context, req dto.DepartmentReq) (dto.DepartmentData, int, error)
	PatchDepartment(ctx context.Context, req dto.DepartmentReq, id, userId string) (dto.DepartmentData, int, error)
	DeleteDepartment(ctx context.Context, user_id string, id string) (dto.DepartmentData, int, error)
}
