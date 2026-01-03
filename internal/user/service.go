package user

import (
	"context"
	"errors"
	"rest-fiber/pkg"

	"gorm.io/gorm"
)

type UserService interface {
	FindAllUsers(ctx context.Context) ([]UserResponseDTO, error)
	FindUserByID(ctx context.Context, id string) (*UserResponseDTO, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) FindAllUsers(ctx context.Context) ([]UserResponseDTO, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	userResponses := make([]UserResponseDTO, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, UserResponseDTO{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	return userResponses, nil
}

func (s *userService) FindUserByID(ctx context.Context, id string) (*UserResponseDTO, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrNotFound
		}
		return nil, err
	}
	return &UserResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}