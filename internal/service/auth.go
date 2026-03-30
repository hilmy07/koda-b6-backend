package service

import (
	"backend/internal/lib"
	"backend/internal/models"
	"backend/internal/repository"
	"errors"
	"time"

	"github.com/matthewhartstonge/argon2"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req models.CreateUserRequest) error {

	existingUser, _ := s.repo.GetByEmail(req.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	argon := argon2.DefaultConfig()

	hash, err := argon.HashEncoded([]byte(req.Password))
	if err != nil {
		return err
	}

	req.Password = string(hash)

	req.Created_at = time.Now()
	req.Updated_at = time.Now()

	return s.repo.CreateUser(req)
}

func (s *AuthService) ProfileRegister(req models.CreateUserRequest) error {

	req.Updated_at = time.Now()

	return s.repo.UpdateUserProfile(req)
}

func (s *AuthService) Login(email string, password string) (*models.User, string, error) {

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}

	ok, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil || !ok {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := lib.GenerateToken(user.Id)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}