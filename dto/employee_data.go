package dto

type EmployeeData struct {
	IdentityNumber   string `json:"identityNumber"`
	Name             string `json:"name"`
	EmployeeImageUri string `json:"employeeImageUri"`
	Gender           string `json:"gender"`
	DepartmentID     string `json:"departmentId"`
	UserId           string `json:"userId"`
}
