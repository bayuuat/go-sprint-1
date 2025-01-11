package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
	executor := d.db.Insert("departments").Rows(goqu.Record{
		"name":          department.Name,
		"user_id":       department.UserId,
	}).Executor()
	fmt.Println(executor.ToSQL())
	_, err := executor.ExecContext(ctx)

	return err
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

func (r *departmentRepository) HasEmployees(ctx context.Context, departmentId string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM employees WHERE department_id = $1"
	err := r.db.QueryRowContext(ctx, query, departmentId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (d departmentRepository) Delete(ctx context.Context, user_id string, id string) error {
	ds := d.db.From("departments").Where(goqu.Ex{
		"department_id": id,
		"user_id":       user_id,
	})

	sql, _, err := ds.Delete().ToSQL()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return fmt.Errorf("Error generating SQL: %w", err)
	}

	_, err = d.db.Exec(sql)
	if err != nil {
		return fmt.Errorf("Error executing SQL: %w", err)
	}

	return err
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
