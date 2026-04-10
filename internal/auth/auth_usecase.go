package auth

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	repo AuthRepository
}

type AuthUseCase interface {
	RegisterUser(ctx context.Context, input *UserDataInput) (string, error)
	AuthUser(ctx context.Context, input *UserDataInput) (string, error)
}

func NewAuthUseCase(repo AuthRepository) AuthUseCase {
	return &authUseCase{
		repo: repo,
	}
}

func (a *authUseCase) RegisterUser(ctx context.Context, input *UserDataInput) (string, error) {
	isUserCreated, err := a.repo.IsUserCreated(ctx, input.Login)
	if err != nil {
		log.Println(err)
		return "", ErrInternalWhileRegisterUser
	}

	if isUserCreated {
		log.Println(err)
		return "", ErrUserAlreayExists
	}

	hashedPassword, err := a.hashPassword(input.Password)
	if err != nil {
		log.Println(err)
		return "", ErrInternalWhileRegisterUser
	}

	err = a.repo.CreateUser(ctx, input.Login, hashedPassword)
	if err != nil {
		log.Println(err)
		return "", ErrInternalWhileRegisterUser
	}

	token, err := a.generateJWTToken(input.Login)
	if err != nil {
		log.Println(err)
		return "", ErrInternalWhileRegisterUser
	}

	return token, nil
}

func (a *authUseCase) AuthUser(ctx context.Context, input *UserDataInput) (string, error) {
	userHashedPassword, err := a.repo.GetUserPasswordHash(ctx, input.Login)
	if err != nil {
		return "", ErrIncorrectUserData //TODO: подумать как сюда интернал ошибку прокинуть лучше
	}

	if !a.checkPasswordHash(input.Password, userHashedPassword) {
		return "", ErrIncorrectUserData
	}

	token, err := a.generateJWTToken(input.Login)
	if err != nil {
		log.Println(err)
		return "", ErrInternalWhileRegisterUser
	}

	return token, nil
}

func (a *authUseCase) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (a *authUseCase) generateJWTToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //TODO: config
		},
		Login: login,
	})

	tokenString, err := token.SignedString([]byte("SECRET_KEY")) //TODO: config
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (a *authUseCase) checkPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err != nil
}
