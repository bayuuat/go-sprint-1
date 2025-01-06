package service

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	cnf            *config.Config
	userRepository domain.UserRepository
}

func NewUser(cnf *config.Config,
	userRepository domain.UserRepository) domain.UserService {
	return &userService{
		cnf:            cnf,
		userRepository: userRepository,
	}
}

func (a userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, int, error) {
	if !req.Action.IsValid() {
		return dto.AuthRes{}, http.StatusBadRequest, domain.ErrInvalidActionItem
	}

	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthRes{}, http.StatusInternalServerError, err
	}

	if req.Action == dto.LoginAction {
		if user.Id == "" {
			return dto.AuthRes{}, http.StatusNotFound, domain.ErrUserNotFound
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return dto.AuthRes{}, http.StatusUnauthorized, domain.ErrInvalidCredential
		}
	} else if req.Action == dto.CreateAction {
		if user.Id != "" {
			return dto.AuthRes{}, http.StatusConflict, domain.ErrEmailExists
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			slog.ErrorContext(ctx, err.Error())
			return dto.AuthRes{}, http.StatusInternalServerError, err
		}

		newUser := domain.User{
			Id:       uuid.New().String(),
			Email:    req.Email,
			Password: string(hashedPassword),
		}

		err = a.userRepository.Save(ctx, &newUser)
		if err != nil {
			slog.ErrorContext(ctx, err.Error())
			return dto.AuthRes{}, http.StatusInternalServerError, err
		}

		user = newUser
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(a.cnf.Secret.Jwt))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthRes{}, http.StatusInternalServerError, err
	}

	var status int
	if req.Action == dto.LoginAction {
		status = http.StatusOK
	} else if req.Action == dto.CreateAction {
		status = http.StatusCreated
	}

	return dto.AuthRes{
		Email:       user.Email,
		AccessToken: tokenString,
	}, status, nil
}

func (a userService) Validate(ctx context.Context, tokenString string) (dto.UserData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cnf.Secret.Jwt), nil
	})
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, err
	}
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return dto.UserData{
				Id:    claims["id"].(string),
				Name:  claims["name"].(string),
				Email: claims["email"].(string),
			}, nil
		}
	}
	return dto.UserData{}, domain.ErrInvalidCredential
}
