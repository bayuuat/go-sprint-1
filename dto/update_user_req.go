package dto

type UpdateUserReq struct {
	Name            *string `json:"name" validate:"omitempty,min=4,max=52"`
	Email           *string `json:"email" validate:"omitempty,email"`
	UserImageUri    *string `json:"userImageUri" validate:"omitempty,uri"`
	CompanyName     *string `json:"companyName" validate:"omitempty,min=4,max=52"`
	CompanyImageUri *string `json:"companyImageUri" validate:"omitempty,uri"`
}
