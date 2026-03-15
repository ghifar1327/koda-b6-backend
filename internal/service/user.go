package service

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"errors"
	"strings"

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

func (s *UserService) ReadAllUser(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUser(ctx)
}
func (s *UserService) ReadByIdUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, req dto.UpdateUsersRequest) error {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Picture) != "" {
		user.Picture.String = req.Picture
	}

	if strings.TrimSpace(req.FullName) != "" {
		user.FullName = req.FullName
	}

	if strings.TrimSpace(req.Email) != "" {
		if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
			return errors.New("invalid email format")
		}
		user.Email = req.Email
	}
	if strings.TrimSpace(req.Password) != "" && len(req.Password) > 5 {
		newPwd, err := utils.HashPassword(req.Password)
		if err != nil {
			return errors.New("password must be at least 5 characters")
		}
		user.Password = newPwd
	}
	if strings.TrimSpace(req.Address) != "" {
		user.Address = req.Address
	}
	if strings.TrimSpace(req.Phone) != "" {
		user.Phone = req.Phone
	}

	if req.RoleId != 0 {
		if req.RoleId != 1 && req.RoleId != 2 {
			return errors.New("invalid role_id")
		}
		user.RoleId = req.RoleId
	}

	return s.repo.UpdateUser(ctx, id, *user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteUser(ctx, id)
}
