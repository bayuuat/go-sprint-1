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
	departmentRepository domain.DepartmentRepository // Add department repository
}

func NewEmployee(cnf *config.Config,
	employeeRepository domain.EmployeeRepository,
	departmentRepository domain.DepartmentRepository) domain.EmployeeService {
	return &employeeService{
		cnf:                  cnf,
		employeeRepository:   employeeRepository,
		departmentRepository: departmentRepository, // Initialize department repository
	}
}

func (ds employeeService) GetEmployees(ctx context.Context, filter dto.EmployeeFilter) ([]dto.EmployeeData, error) {
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

	return employeeData, nil
}

func (ds employeeService) CreateEmployee(ctx context.Context, req dto.EmployeeReq, userId string) (dto.EmployeeData, int, error) {
	departmentIdExists, err := ds.employeeRepository.ExistsDepartmentId(ctx, req.DepartmentID, userId)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, 500, err
	}

	if !departmentIdExists {
		return dto.EmployeeData{}, 400, errors.New(" Department ID not found")
	}

	employee := domain.Employee{
		IdentityNumber:   req.IdentityNumber,
		Name:             req.Name,
		EmployeeImageUri: &req.EmployeeImageUri,
		Gender:           domain.Gender(req.Gender),
		UserId:           userId,
		DepartmentId:     req.DepartmentID,
	}

	err = ds.employeeRepository.Save(ctx, &employee)
	if err != nil {
		return dto.EmployeeData{}, 400, err
	}

	return dto.EmployeeData{
		IdentityNumber:   employee.IdentityNumber,
		Name:             employee.Name,
		EmployeeImageUri: *employee.EmployeeImageUri,
		Gender:           string(employee.Gender),
		DepartmentID:     employee.DepartmentId,
	}, 201, nil
}

func (ds employeeService) PatchEmployee(ctx context.Context, req dto.EmployeeReq, identityNumber, userId string, employeePatch map[string]interface{}) (dto.EmployeeData, int, error) {
	employee, err := ds.employeeRepository.FindById(ctx, identityNumber, userId)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	// IdentityNumber not found in db
	if employee.IdentityNumber == "" {
		return dto.EmployeeData{}, http.StatusNotFound, domain.ErrIdentityNumberNotFound
	}

	if len(employeePatch) == 0 {
		return dto.EmployeeData{
			IdentityNumber:   employee.IdentityNumber,
			Name:             employee.Name,
			EmployeeImageUri: *employee.EmployeeImageUri,
			Gender:           string(employee.Gender),
			DepartmentID:     employee.DepartmentId,
		}, http.StatusOK, nil
	}
	targetIdentityNumber := req.IdentityNumber
	targetEmployee, err := ds.employeeRepository.FindById(ctx, targetIdentityNumber, userId)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	// Target IdentityNumber already taken
	if targetEmployee.IdentityNumber != "" && targetEmployee.IdentityNumber != employee.IdentityNumber {
		return dto.EmployeeData{}, http.StatusConflict, domain.ErrEmployeeExists
	}

	if req.DepartmentID != "" {
		departmentIdExists, err := ds.employeeRepository.ExistsDepartmentId(ctx, req.DepartmentID, userId)

		if err != nil {
			slog.ErrorContext(ctx, err.Error())
			return dto.EmployeeData{}, http.StatusInternalServerError, err
		}

		// Target departmentId not exists
		if !departmentIdExists {
			return dto.EmployeeData{}, http.StatusBadRequest, domain.ErrDepartmentNotFound
		}

		employeePatch["department_id"] = req.DepartmentID
	}

	err = ds.employeeRepository.Update(ctx, userId, employee.IdentityNumber, employeePatch)

	if req.IdentityNumber != "" {
		employee.IdentityNumber = req.IdentityNumber
	}

	if req.Name != "" {
		employee.Name = req.Name
	}

	if req.EmployeeImageUri != "" {
		employee.EmployeeImageUri = &req.EmployeeImageUri
	}

	if req.Gender != "" {
		employee.Gender = domain.Gender(req.Gender)
	}

	if req.DepartmentID != "" {
		employee.DepartmentId = req.DepartmentID
	}

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	return dto.EmployeeData{
		IdentityNumber:   employee.IdentityNumber,
		Name:             employee.Name,
		EmployeeImageUri: *employee.EmployeeImageUri,
		Gender:           string(employee.Gender),
		DepartmentID:     employee.DepartmentId,
	}, http.StatusOK, nil
}

func (ds employeeService) DeleteEmployee(ctx context.Context, user_id string, id string) (dto.EmployeeData, int, error) {
	employee, err := ds.employeeRepository.FindById(ctx, id, user_id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.EmployeeData{}, http.StatusInternalServerError, err
	}

	if employee.IdentityNumber == "" {
		return dto.EmployeeData{}, http.StatusNotFound, domain.ErrEmployeeNotFound
	}

	err = ds.employeeRepository.Delete(ctx, user_id, id)
	if err != nil {
		return dto.EmployeeData{}, 500, err
	}

	return dto.EmployeeData{
		IdentityNumber:   employee.IdentityNumber,
		Name:             employee.Name,
		EmployeeImageUri: *employee.EmployeeImageUri,
		Gender:           string(employee.Gender),
		DepartmentID:     employee.DepartmentId,
	}, http.StatusOK, nil
}

func (ds employeeService) IsEmployeeIDExists(ctx context.Context, identityNumber, userId string) (bool, error) {
	employee, err := ds.employeeRepository.FindById(ctx, identityNumber, userId)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return false, err
	}

	if employee.IdentityNumber != "" {
		return true, nil
	}

	return false, nil
}
