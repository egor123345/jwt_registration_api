package user

import (
	"context"
	"jwt_registration_api/internal/adapters/api/http_handlers/dto"
)

type Service interface {
	Register(ctx context.Context, regInput *dto.RegisterInput) (*dto.RegisterPayload, error)
	Login(ctx context.Context, loginInput *dto.LoginInput) (*dto.LoginPayload, error)
}
