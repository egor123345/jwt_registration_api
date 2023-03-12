package user

import (
	"context"
	"jwt_registration_api/internal/adapters/api/http_handlers/dto"
	"jwt_registration_api/internal/adapters/api/http_handlers/user"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) user.Service {
	return &service{storage: storage}
}

func (s *service) Register(ctx context.Context, regInput *dto.RegisterInput) (*dto.RegisterPayload, error) {
	return nil, nil
}

func (s *service) Login(ctx context.Context, loginInput *dto.LoginInput) (*dto.LoginPayload, error) {
	return nil, nil
}
