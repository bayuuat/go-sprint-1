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
	// Kerjain disini gan
	return domain.ErrInvalidCredential
}

func (d departmentRepository) FindById(ctx context.Context, id string) (domain.Department, error) {
	// Kerjain disini gan
	return domain.Department{}, domain.ErrInvalidCredential
}

func (d departmentRepository) Delete(ctx context.Context, id string) (domain.Department, error) {
	// Kerjain disini gan
	return domain.Department{}, domain.ErrInvalidCredential
}
