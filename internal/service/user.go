package service

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/bayuuat/go-sprint-1/domain"
	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/config"
	"github.com/bayuuat/go-sprint-1/internal/utils"
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

	token, err := utils.GenerateToken(user)

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
		AccessToken: token,
	}, status, nil
}

func (a userService) GetUser(ctx context.Context, email string) (dto.UserData, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, http.StatusInternalServerError, err
	}

	return dto.UserData{
		Id:              user.Id,
		Email:           user.Email,
		Name:            user.Name,
		UserImageUri:    user.UserImageUri,
		CompanyName:     user.CompanyName,
		CompanyImageUri: user.CompanyImageUri,
	}, http.StatusOK, nil
}

func (a userService) PatchUser(ctx context.Context, req dto.UpdateUserReq, email string) (dto.UserData, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, http.StatusInternalServerError, err
	}

	if user.Email == "" {
		return dto.UserData{}, http.StatusNotFound, domain.ErrUserNotFound
	}

	user.Name = req.Name
	user.Email = req.Email
	user.UserImageUri = req.CompanyImageUri
	user.CompanyName = req.CompanyName
	user.CompanyImageUri = req.CompanyImageUri

	err = a.userRepository.Update(ctx, &user)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, http.StatusInternalServerError, err
	}

	return dto.UserData{
		Id:              user.Id,
		Email:           user.Email,
		Name:            user.Name,
		UserImageUri:    user.UserImageUri,
		CompanyName:     user.CompanyName,
		CompanyImageUri: user.CompanyImageUri,
	}, http.StatusOK, nil
}
