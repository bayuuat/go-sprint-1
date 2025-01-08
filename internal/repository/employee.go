package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/domain"
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
	// Kerjain disini gan
	return domain.ErrInvalidCredential
}

func (d employeeRepository) FindById(ctx context.Context, departmentId, identityNumber string) (employee domain.Employee, err error) {
	dataset := d.db.From("employees").Where(goqu.Ex{
		"department_id":   departmentId,
		"identity_number": identityNumber,
	})
	_, err = dataset.ScanStructContext(ctx, &employee)

	return domain.Employee{}, domain.ErrInvalidCredential
}

func (d employeeRepository) ExistsDepartmentId(ctx context.Context, id string) (bool, error) {
	var department domain.Department

	dataset := d.db.From("departments").Where(goqu.Ex{
		"department_id": id,
	})
	_, err := dataset.ScanStructContext(ctx, &department)

	if err != nil {
		return false, err
	}

	departmentIdExists := department.DepartmentId == ""

	return departmentIdExists, nil
}

func (d employeeRepository) Delete(ctx context.Context, id string) (domain.Employee, error) {
	// Kerjain disini gan
	return domain.Employee{}, domain.ErrInvalidCredential
}
