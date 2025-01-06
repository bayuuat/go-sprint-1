package domain

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-1/dto"
)

type User struct {
	Id        string       `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindById(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, int, error)
	Validate(ctx context.Context, tokenString string) (dto.UserData, error)
}
