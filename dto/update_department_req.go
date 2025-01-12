package dto

type UpdateDepartmentReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}
