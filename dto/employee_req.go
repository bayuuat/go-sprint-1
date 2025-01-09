package dto

type EmployeeReq struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,uri"`
	Gender           string `json:"gender" validate:"required,oneof=male female"`
	UserId           string `json:"userId" validate:"required,uuid"`
	DepartmentID     string `json:"departmentId" validate:"required"`
}