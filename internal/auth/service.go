package auth

import (
	"context"
	"errors"
	"fmt"
	"rest-fiber/config"
	"rest-fiber/internal/infra"
	"rest-fiber/internal/user"
	"rest-fiber/pkg"
)

type authService struct {
	userRepo     user.UserRepository
	tokenService infra.TokenService
	emailService infra.EmailService
	env          config.Env
}

func NewAuthService(userRepo user.UserRepository, tokenService infra.TokenService, emailService infra.EmailService, env config.Env) AuthService {
	return &authService{userRepo, tokenService, emailService, env}
}

func (s *authService) Register(ctx context.Context, dto *RegisterRequestDTO) (string, error) {
	exists, err := s.userRepo.FindExistsByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("User Already Exist")
	}
	hashed, err := pkg.HashPassword(dto.Password)
	if err != nil {
		return "", err
	}

	user := &user.User{
		AvatarURL: dto.AvatarURL,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  hashed,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := s.generateVerificationToken(user.ID.String(), user.Email)
	URL := fmt.Sprintf("http://localhost:8080/api/auth?token=%s", token)

	s.emailService.SendEmail(infra.EmailParams{
		Subject: "Verification",
		Message: URL,
		Reciever: infra.EmailReciever{
			Email: user.Email,
		},
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) generateVerificationToken(ID string, Email string) (string, error) {
	return s.tokenService.GenerateToken(&infra.GenerateTokenParams{
		ID:    ID,
		Email: Email,
	}, s.env.JWTVerificationSecret)
}

func (s *authService) VerifyEmailToken(ctx context.Context, token string) error {
	return nil
}
