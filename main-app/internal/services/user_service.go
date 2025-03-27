package services

import (
	"context"

	"github.com/revandpratama/reflect/internal/dto"
	"github.com/revandpratama/reflect/internal/repositories"
	"github.com/revandpratama/reflect/types"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int) (*dto.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) GetUserByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &types.NotFoundError{Message: err.Error()}
		}
		return nil, &types.InternalServerError{Message: err.Error()}
	}

	userResponse := user.ToResponse()

	return &userResponse, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &types.NotFoundError{Message: err.Error()}
		}
		return nil, &types.InternalServerError{Message: err.Error()}
	}

	userResponse := user.ToResponse()

	return &userResponse, nil
}
