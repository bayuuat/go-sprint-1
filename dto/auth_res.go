package dto

type AuthRes struct {
	Email       string `json:"email"`
	AccessToken string `json:"token"`
}
