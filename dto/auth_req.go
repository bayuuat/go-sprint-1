package dto

type AuthReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Action   Action `json:"action"`
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
