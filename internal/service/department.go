package service

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

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

func (ds departmentService) CreateDepartment(ctx context.Context, req dto.DepartmentReq) (dto.DepartmentData, int, error) {
	// Kerjain disini gan
	return dto.DepartmentData{}, 400, errors.New("")
}

func (ds departmentService) PatchDepartment(ctx context.Context, req dto.DepartmentReq, id, userId string) (dto.DepartmentData, int, error) {
	department, err := ds.departmentRepository.FindById(ctx, id, userId)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.DepartmentData{}, http.StatusInternalServerError, err
	}

	if department.DepartmentId == "" {
		return dto.DepartmentData{}, http.StatusNotFound, domain.ErrDepartmentNotFound
	}

	department.Name = req.Name

	err = ds.departmentRepository.Update(ctx, &department)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.DepartmentData{}, http.StatusInternalServerError, err
	}

	return dto.DepartmentData{
		Id:   department.DepartmentId,
		Name: department.Name,
	}, http.StatusOK, nil
}

func (ds departmentService) DeleteDepartment(ctx context.Context, id string) (dto.DepartmentData, int, error) {
	// Kerjain disini gan
	return dto.DepartmentData{}, 400, errors.New("")
}
