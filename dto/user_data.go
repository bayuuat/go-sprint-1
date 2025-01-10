package dto

type UserData struct {
	Id              string  `json:"id"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	UserImageUri    *string `json:"userImageUri"`
	CompanyName     *string `json:"companyName"`
	CompanyImageUri *string `json:"companyImageUri"`
}
