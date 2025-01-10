package domain

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/dto"
)

type Employee struct {
	IdentityNumber   string       `db:"identity_number"`
	Name             string       `db:"name"`
	EmployeeImageUri *string      `db:"employee_image_uri"`
	Gender           Gender       `db:"gender"`
	UserId           string       `db:"user_id"`
	DepartmentId     string       `db:"department_id"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (a Gender) IsValid() bool {
	switch a {
	case Male, Female:
		return true
	}
	return false
}

type EmployeeRepository interface {
	Save(ctx context.Context, employee *Employee) error
	Update(ctx context.Context, employee *Employee) error
	FindById(ctx context.Context, identityNumber, userId string) (Employee, error)
	FindEmployees(ctx context.Context, filter dto.EmployeeFilter) ([]Employee, error)
	ExistsDepartmentId(ctx context.Context, id string, userId string) (bool, error)
	Delete(ctx context.Context, user_id string, id string) error
}

type EmployeeService interface {
	GetEmployees(ctx context.Context, filter dto.EmployeeFilter) ([]dto.EmployeeData, error)
	CreateEmployee(ctx context.Context, req dto.EmployeeReq, id string) (dto.EmployeeData, int, error)
	PatchEmployee(ctx context.Context, req dto.EmployeeReq, identityNumber string, userId string) (dto.EmployeeData, int, error)
	DeleteEmployee(ctx context.Context, user_id string, id string) (dto.EmployeeData, int, error)
	IsEmployeeIDExists(ctx context.Context, identityNumber, userId string) (bool, error)
}
