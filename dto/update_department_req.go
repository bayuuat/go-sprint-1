package dto

type UpdateDepartmentReq struct {
	Name string `json:"name" validate:"omitempty,min=4,max=33"`
}
