package service

import (
	"backend/internal/lib"
	"backend/internal/models"
	"backend/internal/repository"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/matthewhartstonge/argon2"
)

type ForgotPasswordService struct {
	forgotPasswordRepo *repository.ForgotPasswordRepository
	UserRepo *repository.UserRepository
}

func NewForgotPasswordService(UserRepo *repository.UserRepository, forgotPasswordRepo *repository.ForgotPasswordRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		UserRepo: UserRepo, 
		forgotPasswordRepo: forgotPasswordRepo,
	}
}

func GenerateOTP() (string, error) {

	max := big.NewInt(900000)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	code := n.Int64() + 100000

	return fmt.Sprintf("%06d", code), nil
}

func HashPassword(password string) (string, error) {

	hasher := argon2.DefaultConfig()

	hash, err := hasher.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *ForgotPasswordService) RequestForgotPassword(req models.User) error {

	// cek apakah email ada
	existingEmail, err := s.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if existingEmail == nil {
		return errors.New("email not registered")
	}

	// generate OTP
	code, err := GenerateOTP()
	if err != nil {
		return err
	}

	// buat data forgot password
	forgot := models.ForgotPassword{
		Email: req.Email,
		Code:  code,
	}

	fmt.Println(forgot.Code)
	err = lib.SendOTPEmail(req.Email, forgot.Code)
	if err != nil {
		return err
	}

	// simpan ke database
	err = s.forgotPasswordRepo.CreateForgotRequest(forgot)
	if err != nil {
		return err
	}

	return nil
}

func (s *ForgotPasswordService) ResetPassword(reqForgot models.ForgotPassword, reqUser models.User) error {

	// cek apakah email + code valid
	data, err := s.forgotPasswordRepo.GetDataByEmailnCode(reqForgot.Email, reqForgot.Code)
	if err != nil {
		return errors.New("invalid code or email")
	}

	if data == nil {
		return errors.New("code not found")
	}

	// hash password baru
	hashedPassword, err := HashPassword(reqUser.Password)
	if err != nil {
		return err
	}

	// update password user
	err = s.UserRepo.UpdatePasswordByEmail(reqUser.Email, hashedPassword)
	if err != nil {
		return err
	}

	// hapus code setelah dipakai
	err = s.forgotPasswordRepo.DeleteDataByCode(reqForgot.Code)
	if err != nil {
		return err
	}

	return nil
}