package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func validateUser(fullname string, email string, password string) error {
	if len(strings.TrimSpace(fullname)) < 1 {
		return errors.New("Fullname must be at least 1 characters")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("Invalid email format")
	}
	if strings.Index(email, "@") > strings.Index(email, ".") {
		return errors.New("Invalid email domain format")
	}
	if len(password) < 5 {
		return errors.New("Password must be at least 5 characters")
	}
	return nil
}

func (s *UserService) Register(ctx context.Context, req dto.RegisterRequest) error {
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

	return s.repo.Create(ctx, newUser)
}
