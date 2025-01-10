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
	DepartmentId     string       `db:"department_id"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
	UserId           string       `db:"user_id"`
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
	Update(ctx context.Context, userId string, identityNumber string, employeePatch map[string]interface{}) error
	FindById(ctx context.Context, identityNumber, userId string) (Employee, error)
	ExistsDepartmentId(ctx context.Context, id string, userId string) (bool, error)
	Delete(ctx context.Context, id string) (Employee, error)
}

type EmployeeService interface {
	GetEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error)
	CreateEmployee(ctx context.Context, req dto.EmployeeReq, id string) (dto.EmployeeData, int, error)
	PatchEmployee(ctx context.Context, req dto.EmployeeReq, identityNumber, userId string, employee map[string]interface{}) (dto.EmployeeData, int, error)
	DeleteEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error)
}
