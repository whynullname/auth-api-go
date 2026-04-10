package auth

import "context"

type AuthRepository interface {
	IsUserCreated(ctx context.Context, login string) (bool, error)
	CreateUser(ctx context.Context, login string, password string) error
	GetUserPasswordHash(ctx context.Context, login string) (string, error)
}
