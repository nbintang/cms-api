package auth

import "context"

type AuthService interface {
	Register(ctx context.Context, dto *RegisterRequestDTO) (string, error) 
	generateVerificationToken(ID string, Email string) (string, error)
	VerifyEmailToken(ctx context.Context, token string) error
}
