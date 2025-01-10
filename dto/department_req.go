package dto

type DepartmentReq struct {
	// DepartmentId   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	UserId string `json:"userId" validate:"required,uuid"`
}
