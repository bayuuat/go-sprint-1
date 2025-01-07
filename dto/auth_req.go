package dto

type AuthReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Action   Action `json:"action" validate:"required"`
}

type Action string

const (
	CreateAction Action = "create"
	LoginAction  Action = "login"
)

func (a Action) IsValid() bool {
	switch a {
	case CreateAction, LoginAction:
		return true
	}
	return false
}
