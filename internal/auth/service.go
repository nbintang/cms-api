package auth

import (
	"context"
	"rest-fiber/internal/user"
)

type AuthService interface {
	Register(ctx context.Context, dto RegisterRequestDTO) error
}

type authService struct {
	userRepo user.UserRepository
}

func NewAuthService(userRepo user.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s * authService) Register(ctx context.Context, dto RegisterRequestDTO) error {
	var user = user.User{
		Name: dto.Name,
		Email: dto.Email,
		Password: dto.Password,
	}
	
}
