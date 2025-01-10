package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
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
	executor := d.db.Update("departments").Where(goqu.C("department_id").Eq(goqu.L(department.DepartmentId))).Set(department).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

func (d departmentRepository) FindById(ctx context.Context, id, userId string) (department domain.Department, err error) {
	dataset := d.db.From("departments").Where(goqu.Ex{
		"department_id": goqu.L(id),
		"user_id":       userId,
	})
	_, err = dataset.ScanStructContext(ctx, &department)
	return
}

func (d departmentRepository) Delete(ctx context.Context, id string) (domain.Department, error) {
	// Kerjain disini gan
	return domain.Department{}, domain.ErrInvalidCredential
}

func (d departmentRepository) FindAll(ctx context.Context) (result []domain.Department, err error) {
	dataset := d.db.From("departments").Where(goqu.C("created_at").IsNotNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (d departmentRepository) FindByUserId(ctx context.Context, userId string) (result []domain.Department, err error) {
	dataset := d.db.From("departments").Where(goqu.C("user_id").Eq(userId))
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (d departmentRepository) FindAllWithFilter(ctx context.Context, filter *dto.DepartmentFilter) ([]domain.Department, error) {
	query := d.db.From("departments").Where(goqu.Ex{"user_id": filter.UserId})

	if filter.Name != "" {
		query = query.Where(goqu.C("name").ILike("%" + filter.Name + "%"))
	}

	query = query.Limit(uint(filter.Limit)).Offset(uint(filter.Offset))

	var departments []domain.Department
	err := query.ScanStructsContext(ctx, &departments)
	return departments, err
}
