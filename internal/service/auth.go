package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repository.UserRepository
	fgRepo   *repository.AuthRepository
}

func NewAuthService(ur *repository.UserRepository, fg *repository.AuthRepository) *AuthService {
	return &AuthService{
		userRepo: ur,
		fgRepo:   fg,
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
	if err != nil {
		return "", errors.New("Email not Registered")
	}
	valid, err := utils.VerifyPassword(req.Password, user.Password)

	if err != nil {
		return "", err
	}
	if valid {
		token, err := utils.GenerateToken(user.Id)
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
	fmt.Println(otp)
	newData := models.ForgotPassword{
		Email: user.Email,
		Code:  otp,
	}
	return s.fgRepo.CreateForgotPWD(ctx, newData)
}

func (s *AuthService) ResetPassword(ctx context.Context, req dto.ResetPwdRequest) error {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	data, err := s.fgRepo.GetForgotPWDByEmail(ctx, req.Email)
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
	if err := s.userRepo.UpdateUser(ctx, user.Id, *user); err != nil {
		return err
	}
	return s.fgRepo.DeleteForgotPWDByCode(ctx, req.Code)
}
func (s *AuthService) GetUserBYEmail(ctx context.Context, email string) (*models.User , error){
	return  s.userRepo.GetUserByEmail(ctx, email)
}