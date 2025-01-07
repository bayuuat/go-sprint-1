package dto

type UserData struct {
	Id              string  `json:"id"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	UserImageUri    *string `json:"user_image_uri"`
	CompanyName     *string `json:"company_name"`
	CompanyImageUri *string `json:"company_image_uri"`
}
