package service

import (
	"backend/internal/lib"
	"backend/internal/repository"
	"errors"

	"github.com/matthewhartstonge/argon2"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(email string, password string) error {

	existingUser, _ := s.repo.GetByEmail(email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	argon := argon2.DefaultConfig()

	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return err
	}

	err = s.repo.CreateUser(email, string(hash))
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(email string, password string) (string, error) {

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	ok, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil || !ok {
		return "", errors.New("invalid email or password")
	}

	token, err := lib.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}