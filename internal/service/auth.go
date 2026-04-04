package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repository.UserRepository
	authRepo *repository.AuthRepository
}

func NewAuthService(ur *repository.UserRepository, ar *repository.AuthRepository) *AuthService {
	return &AuthService{
		userRepo: ur,
		authRepo: ar,
	}
}

func (a *AuthService) Register(ctx context.Context, req dto.RegisterRequest) error {
	if err := validateUser(req.FullName, req.Email, req.Password); err != nil {
		return err
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := models.User{
		Id:        uuid.New(),
		Email:     req.Email,
		Password:  hash,
		FullName:  req.FullName,
		RoleId:    2,
		Address:   req.Address,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
	}

	return a.userRepo.CreateUser(ctx, newUser)
}

func (a AuthService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	fmt.Println("INPUT PASSWORD:", req.Password)
	fmt.Println("HASH DI DB:", user.Password)
	if err != nil {
		return "", errors.New("Email not Registered")
	}
	valid, err := utils.VerifyPassword(req.Password, user.Password)

	if err != nil {
		return "", err
	}
	if valid {
		token, err := utils.GenerateToken(user)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", errors.New("Invalid Password")
	}
}

func (s *AuthService) RequestForgotPwd(ctx context.Context, email string) error {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	otp, err := utils.GenerateSecureOTP()
	if err != nil {
		return err
	}
	fmt.Printf("name: %s, otp: %d", user.FullName, otp)
	newData := models.ForgotPassword{
		Email: user.Email,
		Code:  otp,
	}
	return s.authRepo.CreateForgotPWD(ctx, newData)
}

func (s *AuthService) ResetPassword(ctx context.Context, req dto.ResetPwdRequest) error {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	data, err := s.authRepo.GetForgotPWDByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if data.Code != req.Code {
		return errors.New("invalid code")
	}

	newPWD, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = newPWD
	if err := s.authRepo.UpdatePassword(ctx, user.Id, user.Password); err != nil {
		return err
	}
	return s.authRepo.DeleteForgotPWDByCode(ctx, req.Code)
}
func (s *AuthService) GetUserBYEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *AuthService) UpdateProfile(ctx context.Context, id uuid.UUID, req dto.UpdateProfileRequest) (models.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	if strings.TrimSpace(req.FullName) != "" {
		user.FullName = req.FullName
	}

	if strings.TrimSpace(req.Email) != "" {
		if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
			return models.User{}, errors.New("invalid email format")
		}
		user.Email = req.Email
	}

	if strings.TrimSpace(req.Address) != "" {
		user.Address = req.Address
	}

	if strings.TrimSpace(req.Phone) != "" {
		user.Phone = req.Phone
	}

	newData, err := s.userRepo.UpdateUser(ctx, id, *user)
	if err != nil {
		return models.User{}, err
	}

	return newData, nil
}

func (s *AuthService) UpdatePicture(ctx context.Context, id uuid.UUID, fileName string) (models.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	if fileName == "" {
		return models.User{}, errors.New("filename is required")
	}

	user.Picture = "/uploads/" + fileName

	newData, err := s.userRepo.UpdateUser(ctx, id, *user)
	if err != nil {
		return models.User{}, err
	}

	return newData, nil
}