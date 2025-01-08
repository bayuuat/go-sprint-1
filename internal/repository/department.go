package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/doug-martin/goqu/v9"
)

type departmentRepository struct {
	db *goqu.Database
}

func NewDepartment(db *sql.DB) domain.DepartmentRepository {
	return &departmentRepository{
		db: goqu.New("default", db),
	}
}

func (d departmentRepository) Save(ctx context.Context, department *domain.Department) error {
	// Kerjain disini gan
	return domain.ErrInvalidCredential
}

func (d departmentRepository) Update(ctx context.Context, department *domain.Department) error {
	executor := d.db.Update("departments").Where(goqu.C("department_id").Eq(department.DepartmentId)).Set(department).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

func (d departmentRepository) FindById(ctx context.Context, id string) (department domain.Department, err error) {
	dataset := d.db.From("departments").Where(goqu.Ex{
		"department_id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &department)
	return
}

func (d departmentRepository) Delete(ctx context.Context, id string) (domain.Department, error) {
	// Kerjain disini gan
	return domain.Department{}, domain.ErrInvalidCredential
}
