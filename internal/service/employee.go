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

func (ds employeeService) PatchEmployee(ctx context.Context, req dto.EmployeeReq, email string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}

func (ds employeeService) DeleteEmployee(ctx context.Context, id string) (dto.EmployeeData, int, error) {
	// Kerjain disini gan
	return dto.EmployeeData{}, 400, errors.New("")
}
