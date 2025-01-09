package service

import (
	"context"
	"errors"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/config"
)

type employeeService struct {
	cnf                *config.Config
	employeeRepository domain.EmployeeRepository
	departmentRepository domain.DepartmentRepository // Add department repository
}

func NewEmployee(cnf *config.Config,
	employeeRepository domain.EmployeeRepository,
	departmentRepository domain.DepartmentRepository) domain.EmployeeService { // Add department repository
	return &employeeService{
		cnf:                cnf,
		employeeRepository: employeeRepository,
		departmentRepository: departmentRepository, // Initialize department repository
	}
}

func (ds employeeService) GetEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}

func (ds employeeService) CreateEmployee(ctx context.Context, req dto.EmployeeReq, email string) (dto.EmployeeData, int, error) {
	// Check if DepartmentID exists
	_, err := ds.departmentRepository.FindById(ctx, req.DepartmentID)
	if err != nil {
		return dto.EmployeeData{}, 400, errors.New("DepartmentID does not exist")
	}

	employee := domain.Employee{
		IdentityNumber: req.IdentityNumber,
		Name: req.Name,
		EmployeeImageUri: &req.EmployeeImageUri,
		Gender: domain.Gender(req.Gender),
		UserId: req.UserId,
		DepartmentId: req.DepartmentID,
	}

	err = ds.employeeRepository.Save(ctx, &employee)
	if err != nil {
		return dto.EmployeeData{}, 400, err
	}

	return dto.EmployeeData{
		IdentityNumber: employee.IdentityNumber,
		Name: employee.Name,
		EmployeeImageUri: *employee.EmployeeImageUri,
		Gender: string(employee.Gender),
		DepartmentID: employee.DepartmentId,
	}, 201, nil
}

func (ds employeeService) PatchEmployee(ctx context.Context, req dto.EmployeeReq, email string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}

func (ds employeeService) DeleteEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}
