package dto

type DepartmentReq struct {
	// DepartmentId   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	UserId string `json:"userId" validate:"required,uuid"`
}

type DepartmentFilter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}
