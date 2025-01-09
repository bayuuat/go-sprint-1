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
	executor := d.db.Insert("employees").Rows(employee).Executor()
	_, err := executor.ExecContext(ctx)

	return err

}

func (d employeeRepository) Update(ctx context.Context, employee *domain.Employee) error {
	// Kerjain disini gan
	return domain.ErrInvalidCredential
}

func (d employeeRepository) FindById(ctx context.Context, id string) (domain.Employee, error) {
	// Kerjain disini gan
	return domain.Employee{}, domain.ErrInvalidCredential
}

func (d employeeRepository) Delete(ctx context.Context, id string) (domain.Employee, error) {
	// Kerjain disini gan
	return domain.Employee{}, domain.ErrInvalidCredential
}
