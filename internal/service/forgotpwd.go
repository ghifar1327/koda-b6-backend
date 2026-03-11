package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
)

type ForgotPwdService struct {
	userRepo *repository.UserRepository
	fgRepo   *repository.ForgotPWDRepository
}

func NewForgotPwdService(ur *repository.UserRepository, fg *repository.ForgotPWDRepository) *ForgotPwdService {
	return &ForgotPwdService{
		userRepo: ur,
		fgRepo:   fg,
	}
}

func (s *ForgotPwdService) RequestForgotPwd(ctx context.Context, email string) error {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	otp, err := utils.GenerateSecureOTP()
	if err != nil {
		return err
	}
	newData := models.ForgotPassword{
		Email: user.Email,
		Code:  otp,
	}
	return s.fgRepo.CreateForgotPWD(ctx, newData)
}

func (s *ForgotPwdService) ResetPassword(ctx context.Context, req dto.ResetPwdRequest) error {
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
