package dto

type DepartmentReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}

type DepartmentFilter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}
