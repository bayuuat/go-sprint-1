package dto

type EmployeeReq struct {
	IdentityNumber   string `json:"identityNumber" validate:"required,min=5,max=33"`
	Name             string `json:"name" validate:"required,min=4,max=33"`
	EmployeeImageUri string `json:"employeeImageUri" validate:"required,uri"`
	Gender           string `json:"gender" validate:"required,oneof=male female"`
	DepartmentID     string `json:"departmentId" validate:"required"`
}

type EmployeeFilter struct {
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Name           string `json:"name"`
	UserId         string `json:"user_id"`
	IdentityNumber string `json:"identity_number"`
	Gender         string `json:"gender"`
	DepartmentID   string `json:"department_id"`
}
