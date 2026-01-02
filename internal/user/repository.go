package user

import (
	"context"

	"gorm.io/gorm"
)


type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll(ctx context.Context) ([]User, error) {
	var user []User
	if err := r.db.WithContext(ctx).Find(&user).Error; err != nil {
		return nil , err
	};
	return user, nil
}

