package service

import (
	"backend/internal/dto"
	"backend/internal/repository"
	"context"
	"errors"
	"strings"
)

type MasterService struct {
	repo *repository.MasterRepository
}

func NewMasterService(repo *repository.MasterRepository) *MasterService {
	return &MasterService{repo: repo}
}

func (s *MasterService) Create(ctx context.Context, table string, req dto.CreateMasterRequest) error {
	return s.repo.Create(ctx, table, req)
}

func (s *MasterService) GetAll(ctx context.Context, table string) ([]dto.Master, error) {
	return s.repo.GetAll(ctx, table)
}

func (s *MasterService) GetById(ctx context.Context, table string, id int) (dto.Master, error) {
	return s.repo.GetById(ctx, table, id)
}

func (s *MasterService) Update(ctx context.Context, table string, id int, req dto.UpdateMasterRequest) error {
	data, err := s.repo.GetById(ctx, table, id)
	if err != nil {
		return err
	}

	if req.Name == nil && req.AddPrice == nil {
		return errors.New("no fields to update")
	}

	if strings.TrimSpace(*req.Name) != "" {
		data.Name = *req.Name
	}

	if *req.AddPrice > 0 {
		data.AddPrice = req.AddPrice
	}else{
		return errors.New("add price cannot be minus")
	}

	newData := dto.UpdateMasterRequest{
		Name:     &data.Name,
		AddPrice: data.AddPrice,
	}

	return s.repo.Update(ctx, table, id, newData)
}

func (s *MasterService) Delete(ctx context.Context, table string, id int) error {
	return s.repo.Delete(ctx, table, id)
}
