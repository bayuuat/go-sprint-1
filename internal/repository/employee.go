package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/doug-martin/goqu/v9"
)

type employeeRepository struct {
	db *goqu.Database
}

func NewEmployee(db *sql.DB) domain.EmployeeRepository {
	return &employeeRepository{
		db: goqu.New("default", db),
	}
}

func (d employeeRepository) Save(ctx context.Context, employee *domain.Employee) error {
	// Kerjain disini gan
	return domain.ErrInvalidCredential
}

func (d employeeRepository) Update(ctx context.Context, employee *domain.Employee) error {
	executor := d.db.Update("employees").Where(goqu.Ex{
		"user_id":         employee.UserId,
		"identity_number": employee.IdentityNumber,
	}).Set(employee).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (d employeeRepository) FindById(ctx context.Context, identityNumber, userId string) (employee domain.Employee, err error) {
	dataset := d.db.From("employees").Where(goqu.Ex{
		"user_id":         userId,
		"identity_number": identityNumber,
	})
	_, err = dataset.ScanStructContext(ctx, &employee)

	return employee, nil
}

func (d employeeRepository) ExistsDepartmentId(ctx context.Context, id, userId string) (bool, error) {
	var department domain.Department

	dataset := d.db.From("departments").Where(goqu.Ex{
		"department_id": goqu.L(id),
		"user_id":       userId,
	})
	_, err := dataset.ScanStructContext(ctx, &department)

	if err != nil {
		return false, err
	}

	departmentIdExists := department.DepartmentId != ""

	return departmentIdExists, nil
}

func (d employeeRepository) Delete(ctx context.Context, id string) (domain.Employee, error) {
	// Kerjain disini gan
	return domain.Employee{}, domain.ErrInvalidCredential
}

func (d employeeRepository) FindByUserId(ctx context.Context, userId string) (result []domain.Employee, err error) {
	dataset := d.db.From("employees").Where(goqu.C("user_id").Eq(userId))
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (d employeeRepository) FindEmployees(
	ctx context.Context,
	filter dto.EmployeeFilter) ([]domain.Employee, error) {
	dataset := d.db.From("employees").Where(goqu.C("user_id").Eq(filter.UserId))

	if filter.IdentityNumber != "" {
		dataset = dataset.Where(goqu.C("identity_number").ILike(filter.IdentityNumber))
	}

	if filter.Name != "" {
		dataset = dataset.Where(goqu.C("name").ILike("%" + filter.Name + "%"))
	}

	if filter.Gender != "" {
		if filter.Gender == "male" || filter.Gender == "female" {
			dataset = dataset.Where(goqu.C("gender").Eq(filter.Gender))
		} else {
			return []domain.Employee{}, nil
		}
	}

	if filter.DepartmentID != "" {
		dataset = dataset.Where(goqu.C("department_id").Eq(filter.DepartmentID))
	}

	if filter.Limit <= 0 {
		filter.Limit = 5
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}
	dataset = dataset.Limit(uint(filter.Limit)).Offset(uint(filter.Offset))

	var employees []domain.Employee
	err := dataset.ScanStructsContext(ctx, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}