package dto

type UpdateUserReq struct {
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	UserImageUri    *string `json:"user_image_uri"`
	CompanyName     *string `json:"company_name"`
	CompanyImageUri *string `json:"company_image_uri"`
}
