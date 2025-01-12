package dto

type UpdateUserReq struct {
	Name            *string `json:"name" validate:"required,min=4,max=52"`
	Email           *string `json:"email" validate:"required,email"`
	UserImageUri    *string `json:"userImageUri" validate:"required,uri,accessibleuri"`
	CompanyName     *string `json:"companyName" validate:"required,min=4,max=52"`
	CompanyImageUri *string `json:"companyImageUri" validate:"required,uri,accessibleuri"`
}
