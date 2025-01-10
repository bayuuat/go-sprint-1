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

func (ds departmentService) GetDepartmentsWithFilter(ctx context.Context, filter dto.DepartmentFilter) ([]dto.DepartmentData, int, error) {
	if filter.Limit <= 0 {
		filter.Limit = 5
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}

	departments, err := ds.departmentRepository.FindAllWithFilter(ctx, &filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(departments) == 0 {
		return []dto.DepartmentData{}, http.StatusOK, nil
	}

	var departmentData []dto.DepartmentData
	for _, v := range departments {
		departmentData = append(departmentData, dto.DepartmentData{
			Id:   v.DepartmentId,
			Name: v.Name,
		})
	}

	return departmentData, http.StatusOK, nil
}

func (ds departmentService) CreateDepartment(ctx context.Context, req dto.DepartmentReq) (dto.DepartmentData, int, error) {
	// Kerjain disini gan
	return dto.DepartmentData{}, 400, errors.New("")
}

func (ds departmentService) PatchDepartment(ctx context.Context, req dto.UpdateDepartmentReq, id, userId string) (dto.DepartmentData, int, error) {
	department, err := ds.departmentRepository.FindById(ctx, id, userId)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.DepartmentData{}, http.StatusInternalServerError, err
	}

	if department.DepartmentId == "" {
		return dto.DepartmentData{}, http.StatusNotFound, domain.ErrDepartmentNotFound
	}

	if req.Name == "" {
		return dto.DepartmentData{
			Id:   department.DepartmentId,
			Name: department.Name,
		}, http.StatusOK, nil
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

func (ds departmentService) DeleteDepartment(ctx context.Context, user_id string, id string) (dto.DepartmentData, int, error) {
	department, err := ds.departmentRepository.FindById(ctx, id, user_id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.DepartmentData{}, http.StatusInternalServerError, err
	}

	if department.DepartmentId == "" {
		return dto.DepartmentData{}, http.StatusNotFound, domain.ErrDepartmentNotFound
	}

	hasEmployees, err := ds.departmentRepository.HasEmployees(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.DepartmentData{}, http.StatusInternalServerError, err
	}

	if hasEmployees {
		return dto.DepartmentData{}, http.StatusConflict, domain.ErrDepartmentHasEmployees
	}

	err = ds.departmentRepository.Delete(ctx, user_id, id)
	if err != nil {
		return dto.DepartmentData{}, 500, err
	}

	return dto.DepartmentData{
		Id:   department.DepartmentId,
		Name: department.Name,
	}, http.StatusOK, nil
}
