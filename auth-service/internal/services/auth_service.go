package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/revandpratama/reflect/auth-service/helper"
	"github.com/revandpratama/reflect/auth-service/internal/dto"
	"github.com/revandpratama/reflect/auth-service/internal/entities"
	"github.com/revandpratama/reflect/auth-service/internal/repositories"
)

type authService struct {
	repository repositories.AuthRepository
}

type AuthService interface {
	Login(ctx context.Context, email, password string) error
	Register(ctx context.Context, req dto.RegisterRequest) error
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

func (s *authService) Login(ctx context.Context, email, password string) error {

	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	if err := helper.ValidatePassword(user.Password, password); err != nil {
		return fmt.Errorf("login failed: %v", err)
	}

	return nil
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) error {

	if s.repository.IsEmailExists(ctx, req.Email) {
		return errors.New("unable to create new user: email already exists")
	}

	encryptedPasswd, err := helper.EncryptPassword(req.Password)
	if err != nil {
		return fmt.Errorf("unable to create new user: %v", err)
	}

	user := entities.User{
		RoleID:   1,
		Name:     req.Name,
		Email:    req.Email,
		Password: encryptedPasswd,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("unable to create new user: %v", err)
	}

	return nil
}
