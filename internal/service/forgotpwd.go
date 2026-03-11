package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
)

type ForgotPwdService struct {
	userRepo *repository.UserRepository
	fgRepo   *repository.DataRepository
}

func newForgotPwdService(ur *repository.UserRepository, fg *repository.DataRepository) *ForgotPwdService {
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
	return s.fgRepo.CreateData(ctx, newData)
}


func (s *ForgotPwdService) ResetPassword(ctx context.Context, email string, code int) error {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	
	data, err := s.fgRepo.GetDataByEmail(ctx, email)
	if err != nil {
		return err
	}

	if data.Code != code {
		return errors.New("invalid code")
	}
	if err := s.userRepo.UpdateUser(ctx , user.Id , *user); err !=nil{
		return err
	}
	return s.fgRepo.DeleteDataByCode(ctx, code)
}