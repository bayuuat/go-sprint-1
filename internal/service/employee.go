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
	cnf                  *config.Config
	employeeRepository   domain.EmployeeRepository
	departmentRepository domain.DepartmentRepository
}

func NewEmployee(cnf *config.Config,
	employeeRepository domain.EmployeeRepository,
	departmentRepository domain.DepartmentRepository) domain.EmployeeService {
	return &employeeService{
		cnf:                  cnf,
		employeeRepository:   employeeRepository,
		departmentRepository: departmentRepository,
	}
}

// func (ds employeeService) GetEmployeeByUserId(ctx context.Context, userId string) ([]dto.EmployeeData, int, error) {
// 	employee, err := ds.employeeRepository.FindByUserId(ctx, userId)
// 	if err != nil {
// 		slog.ErrorContext(ctx, err.Error())
// 		return nil, http.StatusInternalServerError, err
// 	}
// 	var employeeData []dto.EmployeeData
// 	for _, v := range employee {
// 		employeeData = append(employeeData, dto.EmployeeData{
// 			IdentityNumber:   v.IdentityNumber,
// 			Name:             v.Name,
// 			EmployeeImageUri: *v.EmployeeImageUri,
// 			DepartmentID:     v.DepartmentId,
// 		})
// 	}
// 	return employeeData, http.StatusOK, nil
// }

func (ds employeeService) GetEmployees(ctx context.Context, filter dto.EmployeeFilter) ([]domain.Employee, error) {
	if filter.Limit <= 0 {
		filter.Limit = 5
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}

	employees, err := ds.employeeRepository.FindEmployees(ctx, filter)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return nil, err
	}

	var employeeData []dto.EmployeeData
	for _, v := range employees {
		employeeData = append(employeeData, dto.EmployeeData{
			IdentityNumber:   v.IdentityNumber,
			Name:             v.Name,
			EmployeeImageUri: *v.EmployeeImageUri,
			Gender:           string(v.Gender),
			DepartmentID:     v.DepartmentId,
			UserId:           v.UserId,
		})
	}

	var employeesDomain []domain.Employee
	for _, v := range employees {
		employeesDomain = append(employeesDomain, domain.Employee{
			IdentityNumber:   v.IdentityNumber,
			Name:             v.Name,
			EmployeeImageUri: v.EmployeeImageUri,
			Gender:           v.Gender,
			DepartmentId:     v.DepartmentId,
			UserId:           v.UserId,
		})
	}

	return employeesDomain, nil
}

func (ds employeeService) CreateEmployee(ctx context.Context, req dto.EmployeeReq, email string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}

func (ds employeeService) PatchEmployee(ctx context.Context, req dto.EmployeeReq, identityNumber, userId string) (dto.EmployeeData, int, error) {
	employee, err := ds.employeeRepository.FindById(ctx, identityNumber, userId)

	targetIdentityNumber := req.IdentityNumber
	targetEmployee, err := ds.employeeRepository.FindById(ctx, targetIdentityNumber, userId)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	// IdentityNumber params not found
	if employee.IdentityNumber == "" {
		return dto.EmployeeData{}, http.StatusNotFound, domain.ErrIdentityNumberNotFound
	}

	// Target IdentityNumber already taken
	if targetEmployee.IdentityNumber != "" && targetEmployee.IdentityNumber != employee.IdentityNumber {
		return dto.EmployeeData{}, http.StatusConflict, domain.ErrEmployeeExists
	}

	departmentIdExists, err := ds.employeeRepository.ExistsDepartmentId(ctx, req.DepartmentID, userId)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	// Target departmentId not exists
	if !departmentIdExists {
		return dto.EmployeeData{}, http.StatusBadRequest, domain.ErrDepartmentNotFound
	}

	employee.IdentityNumber = req.IdentityNumber
	employee.Name = req.Name
	employee.EmployeeImageUri = &req.EmployeeImageUri
	employee.Gender = domain.Gender(req.Gender)
	employee.DepartmentId = req.DepartmentID

	err = ds.employeeRepository.Update(ctx, &employee)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	return dto.EmployeeData{
		IdentityNumber:   employee.DepartmentId,
		Name:             employee.Name,
		EmployeeImageUri: *employee.EmployeeImageUri,
		Gender:           string(employee.Gender),
		DepartmentID:     employee.DepartmentId,
	}, http.StatusOK, nil
}

func (ds employeeService) DeleteEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}
