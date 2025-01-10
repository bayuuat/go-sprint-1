package service

import (
	"context"
	"errors"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/config"
)

type departmentService struct {
	cnf                  *config.Config
	departmentRepository domain.DepartmentRepository
}

func NewDepartment(cnf *config.Config,
	departmentRepository domain.DepartmentRepository) domain.DepartmentService {
	return &departmentService{
		cnf:                  cnf,
		departmentRepository: departmentRepository,
	}
}

func (ds departmentService) GetDepartment(ctx context.Context, id string) (dto.DepartmentData, int, error) {
	// Kerjain disini gan
	return dto.DepartmentData{}, 400, errors.New("")
}

func (ds departmentService) CreateDepartment(ctx context.Context, req dto.DepartmentReq, email string) (dto.DepartmentData, int, error) {
	department := domain.Department{
		Name: req.Name,
		UserId: req.UserId,
	}

	err := ds.departmentRepository.Save(ctx, &department)
	if err != nil {
		return dto.DepartmentData{}, 400, err
	}

	return dto.DepartmentData{
		DepartmentId: department.DepartmentId,
		Name: department.Name,
		UserId: department.UserId,
	}, 201, nil
}

func (ds departmentService) PatchDepartment(ctx context.Context, req dto.DepartmentReq, email string) (dto.DepartmentData, int, error) {
	// Kerjain disini gan
	return dto.DepartmentData{}, 400, errors.New("")
}

func (ds departmentService) DeleteDepartment(ctx context.Context, id string) (dto.DepartmentData, int, error) {
	// Kerjain disini gan
	return dto.DepartmentData{}, 400, errors.New("")
}
