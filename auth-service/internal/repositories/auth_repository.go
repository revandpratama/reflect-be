package repositories

import (
	"context"

	"github.com/revandpratama/reflect/auth-service/internal/entities"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) error
	IsEmailExists(ctx context.Context, email string) bool
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User

	err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error

	return user, err
}

func (r *authRepository) IsEmailExists(ctx context.Context, email string) bool {
	var user []entities.User

	r.db.WithContext(ctx).Find(&user, "email = ?", email)

	return len(user) > 0
}

func (r *authRepository) CreateUser(ctx context.Context, user entities.User) error {
	err := r.db.WithContext(ctx).Create(&user).Error

	return err
}
