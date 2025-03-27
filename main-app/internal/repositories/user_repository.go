package repositories

import (
	"context"

	"github.com/revandpratama/reflect/internal/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	GetUserByID(ctx context.Context, id int) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User

	err := r.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error

	return user, err
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entities.User, error) {
	var user entities.User

	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&user).Error

	return user, err
}