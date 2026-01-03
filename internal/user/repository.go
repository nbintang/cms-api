package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	FindExistsByEmail(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, dto *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(ctx context.Context) ([]User, error) {
	var user []User
	err := r.db.WithContext(ctx).Find(&user).Error
	return user, err
}

func (r *userRepository) FindByID(ctx context.Context, id string) (User, error) {
	var user User
	err := r.db.WithContext(ctx).Scopes(WhereID(id), SelectPublicFields).Take(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.WithContext(ctx).Scopes(WhereEmail(email), SelectPublicFields).Take(&user).Error
	return user, err
}

func (r *userRepository) FindExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&User{}).Scopes(WhereEmail(email)).Count(&count).Error
	return count > 0, err
}

func (r *userRepository) Create(ctx context.Context, dto *User) error {
	dto.Role = Member
	err := r.db.WithContext(ctx).Create(&dto).Error
	return err
}
