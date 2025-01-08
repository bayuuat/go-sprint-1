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

type employeeService struct {
	cnf                *config.Config
	employeeRepository domain.EmployeeRepository
}

func NewEmployee(cnf *config.Config,
	employeeRepository domain.EmployeeRepository) domain.EmployeeService {
	return &employeeService{
		cnf:                cnf,
		employeeRepository: employeeRepository,
	}
}

func (ds employeeService) GetEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}

func (ds employeeService) CreateEmployee(ctx context.Context, req dto.EmployeeReq, email string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}

func (ds employeeService) PatchEmployee(ctx context.Context, req dto.EmployeeReq, identityNumber string) (dto.EmployeeData, int, error) {
	employee, err := ds.employeeRepository.FindById(ctx, identityNumber)

	targetId := req.IdentityNumber
	targetEmployee, err := ds.employeeRepository.FindById(ctx, targetId)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	if employee.IdentityNumber == "" {
		return dto.EmployeeData{}, http.StatusNotFound, domain.ErrDepartmentNotFound
	}

	if targetEmployee.IdentityNumber != "" {
		return dto.EmployeeData{}, http.StatusConflict, domain.ErrEmployeeExists
	}

	departmentIdExists, err := ds.employeeRepository.ExistsDepartmentId(ctx, req.DepartmentID)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	if !departmentIdExists {
		return dto.EmployeeData{}, http.StatusBadRequest, domain.ErrDepartmentNotFound
	}

	employee.Name = req.Name

	err = ds.employeeRepository.Update(ctx, &employee)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	return dto.EmployeeData{
		Id:   employee.DepartmentId,
		Name: employee.Name,
	}, http.StatusOK, nil
}

func (ds employeeService) DeleteEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}
